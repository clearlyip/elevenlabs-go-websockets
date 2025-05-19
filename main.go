// Client implementation for Eleven Labs TTS API:
// 		- Streaming Websocket: wss://api.elevenlabs.io/v1/text-to-speech/:voice_id/stream-input
// 		- Multi-Context Websocket: wss://api.elevenlabs.io/v1/text-to-speech/:voice_id/multi-stream-input

package elevenlabs

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/nrednav/cuid2"
)

const JSON_CONTENT_TYPE = "application/json"
const ELEVEN_BASEURL_HTTPS = "https://api.elevenlabs.io/v1"
const ELEVEN_BASEURL_WSS = "wss://api.elevenlabs.io/v1"
const MULTI_CONTEXT_MAX_REQUESTS = 5
const STD_MAX_REQUESTS = 1

type Client struct {
	apiKey  string
	timeout time.Duration
	ctx     context.Context
}

type MultiClient struct {
	apiKey                   string
	timeout                  time.Duration
	ctx                      context.Context
	activeRequests           map[string]struct{} // Active multi-context requests
	cmu                      sync.RWMutex
	TextReader               chan string
	AlignmentResponseChannel chan StreamingOutputMultiCtxResponse
}

type VoiceSettings struct {
	SimilarityBoost float32 `json:"similarity_boost"`
	Stability       float32 `json:"stability"`
	Style           float32 `json:"style,omitempty"`
	SpeakerBoost    bool    `json:"use_speaker_boost,omitempty"`
}

type GenerationConfig struct {
	ChunkLengthSchedule []int `json:"chunk_length_schedule"`
}

type QueryFunc func(*neturl.Values)

type StreamingInputResponse struct {
	Audio               string                    `json:"audio"`
	IsFinal             bool                      `json:"isFinal"`
	NormalizedAlignment StreamingAlignmentSegment `json:"normalizedAlignment"`
	Alignment           StreamingAlignmentSegment `json:"alignment"`
}

type StreamingOutputResponse struct {
	IsFinal             bool                      `json:"isFinal"`
	NormalizedAlignment StreamingAlignmentSegment `json:"normalizedAlignment"`
	Alignment           StreamingAlignmentSegment `json:"alignment"`
}

type StreamingOutputMultiCtxRawResponse struct {
	IsFinal             bool                      `json:"isFinal"`
	NormalizedAlignment StreamingAlignmentSegment `json:"normalizedAlignment"`
	Alignment           StreamingAlignmentSegment `json:"alignment"`
	ContextId           string                    `json:"contextId"`
	ContextIdAlt        string                    `json:"context_id"`
}

type StreamingOutputMultiCtxResponse struct {
	IsFinal             bool                      `json:"isFinal"`
	NormalizedAlignment StreamingAlignmentSegment `json:"normalizedAlignment"`
	Alignment           StreamingAlignmentSegment `json:"alignment"`
	ContextId           string                    `json:"contextId"`
}

type TextToSpeechInputStreamingRequest struct {
	Text             string            `json:"text"`
	ContextID        string            `json:"context_id,omitempty"`
	VoiceSettings    *VoiceSettings    `json:"voice_settings,omitempty"`
	GenerationConfig *GenerationConfig `json:"generation_config,omitempty"`
}

type WsStreamingOutputChannel chan StreamingOutputResponse

// Standard Websocket Client
func NewClient(ctx context.Context, apiKey string, reqTimeout time.Duration) *Client {
	return &Client{apiKey: apiKey, timeout: reqTimeout, ctx: ctx}
}

// Multi-Context Websocket Session
func NewMultiContextSession(ctx context.Context, apiKey string, reqTimeout time.Duration, TextReader chan string, AlignmentResponseChannel chan StreamingOutputMultiCtxResponse, AudioResponsePipe io.Writer, voiceID string, modelID string, req TextToSpeechInputMultiStreamingRequest, queries ...QueryFunc) error {

	c := &MultiClient{
		apiKey:                   apiKey,
		timeout:                  reqTimeout,
		ctx:                      ctx,
		activeRequests:           make(map[string]struct{}),
		TextReader:               TextReader,
		AlignmentResponseChannel: AlignmentResponseChannel,
	}

	err := c.MultiCtxStreamingRequest(TextReader, AlignmentResponseChannel, AudioResponsePipe, voiceID, modelID, queries...)
	if err != nil {
		return err
	}
	return nil
}

func GetUserCapacity(apiKey string) (*UserAndCapacity, error) {
	u, err := User(apiKey)
	if err != nil {
		return nil, err
	}
	sub := u.Subscription

	return &UserAndCapacity{
		UserID:       u.UserID,
		Subscription: u.Subscription,
		HasCapacity: sub.CharacterCount <= sub.CharacterLimit ||
			sub.CanExtendCharacterLimit,
	}, nil
}

// API Get user
func User(apiKey string) (*UserData, error) {
	url := fmt.Sprintf("%s/user", ELEVEN_BASEURL_HTTPS)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("xi-api-key", apiKey)

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r UserData
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func GetVoice(apiKey string, voiceId string) (*GetVoiceVoice, error) {
	url := fmt.Sprintf("%s/voices/%s", ELEVEN_BASEURL_HTTPS, voiceId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("xi-api-key", apiKey)

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r GetVoiceVoice
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func SharedVoices(apiKey string, params ListVoicesParams) (*ListVoicesResponse, error) {
	url := fmt.Sprintf("%s/shared-voices", ELEVEN_BASEURL_HTTPS)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("xi-api-key", apiKey)

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r ListVoicesResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
	//https://api.elevenlabs.io/v1/shared-voices
}

func ValidateLanguageAndModel(apiKey string, voiceId string, modelName string) (bool, error) {
	gv, err := GetVoice(apiKey, voiceId) // Voice exists?
	if err != nil {
		return false, fmt.Errorf("voice lookup failed")
	}
	if gv.VoiceID == "" {
		return false, fmt.Errorf("voice not found")
	}

	// Does voice support the model?
	voiceSupportsModel := false
	for _, s := range gv.HighQualityBaseModelIDs {
		if s == modelName {
			voiceSupportsModel = true
		}
	}
	if !voiceSupportsModel {
		return false, fmt.Errorf("voice does not support the model specified")
	}

	return true, nil
}

// Standard Websocket Request
func (c *Client) StreamingRequest(TextReader chan string, AlignmentResponseChannel chan StreamingOutputResponse, AudioResponsePipe io.Writer, voiceID string, modelID string, req TextToSpeechInputStreamingRequest, queries ...QueryFunc) error {
	driverActive := true // Driver shut down?
	driverError := false // Unexpected errors

	// url := fmt.Sprintf("%s/text-to-speech/%s/stream-input?model_id=%s", ELEVEN_BASEURL_WSS, voiceID, modelID)
	url := fmt.Sprintf("%s/text-to-speech/%s/multi-stream-input?model_id=%s&inactivity_timeout=180&sync_alignment=true", ELEVEN_BASEURL_WSS, voiceID, modelID)
	multiCtx := cuid2.Generate()

	headers := http.Header{}
	headers.Add("Accept", "*/*")
	headers.Add("Content-Type", JSON_CONTENT_TYPE)
	if c.apiKey != "" {
		headers.Add("xi-api-key", c.apiKey)
	}

	u, err := neturl.Parse(url)
	if err != nil {
		return err
	}

	q := u.Query()
	for _, qf := range queries {
		qf(&q)
	}
	u.RawQuery = q.Encode()

	conn, _, err := websocket.DefaultDialer.DialContext(c.ctx, u.String(), headers)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("🎲🎲🎲 Connected to Eleven Labs TTS WebSocket")

	initReq := TextToSpeechInputMultiStreamingRequest{
		Text:      " ",
		ContextID: multiCtx,
	}

	fmt.Println("🎲🎲🎲 Sending initialization request", "initReq", initReq)

	// Send initial request
	if err := conn.WriteJSON(initReq); err != nil {
		return err
	}

	// Input watcher
	inputCtx, inputCancel := context.WithCancel(context.Background())

	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	// Response watching
	go func(wg *sync.WaitGroup, errCh chan<- error) {
		defer wg.Done()
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				if !driverActive {
					return
				}
				var input StreamingInputResponse
				var response StreamingOutputResponse
				if err := conn.ReadJSON(&input); err != nil {
					if driverActive {
						errCh <- err
						driverError = true
						inputCancel()
					}
					return
				}

				b, err := base64.StdEncoding.DecodeString(input.Audio)
				if err != nil {
					if driverActive {
						errCh <- err
						driverError = true
						inputCancel()
					}
					return
				}
				// Send audio through the pipeline
				if _, err := AudioResponsePipe.Write(b); err != nil {
					break
				}

				// Send non-audio via the response channel
				response = StreamingOutputResponse{
					IsFinal:             input.IsFinal,
					NormalizedAlignment: input.NormalizedAlignment,
					Alignment:           input.Alignment,
				}
				AlignmentResponseChannel <- response
			}
		}
	}(&wg, errCh)

	// Input watching
InputWatcher:
	for {
		select {
		case <-inputCtx.Done():
			driverActive = false
			break InputWatcher
		case <-c.ctx.Done():
			driverActive = false
			break InputWatcher
		case chunk, ok := <-TextReader:
			if !ok || !driverActive {
				break InputWatcher
			}
			ch := &TextToSpeechInputMultiStreamingRequest{Text: chunk, ContextID: multiCtx}
			if err := conn.WriteJSON(ch); err != nil {
				errCh <- err
				break InputWatcher
			}
		}
	}

	// Send final "" to close out TTS buffer
	if driverActive && !driverError {
		if err := conn.WriteJSON(map[string]string{"text": ""}); err != nil {
			if c.ctx.Err() == nil {
				errCh <- err
			}
		}
	}
	conn.Close()
	wg.Wait()

	// Errors?
	select {
	case readErr := <-errCh:
		if driverActive || driverError {
			// Only send if the driver is active or the unexpected error flag is active
			return readErr
		} else {
			return nil
		}
	default:
	}

	return nil
}

func (c *MultiClient) activeContextCount() int {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return len(c.activeRequests)
}

func (c *MultiClient) hasMultiCtx(id string) bool {
	c.cmu.RLock()
	_, ok := c.activeRequests[id]
	c.cmu.RUnlock()
	return ok
}

func (c *MultiClient) addMultiCtx(id string) {
	c.cmu.Lock()
	c.activeRequests[id] = struct{}{}
	c.cmu.Unlock()
}

func (c *MultiClient) removeMultiCtx(id string) {
	c.cmu.Lock()
	delete(c.activeRequests, id)
	c.cmu.Unlock()
}

func (c *MultiClient) HasCapacity() bool {
	c.cmu.RLock()
	defer c.cmu.RUnlock()
	return len(c.activeRequests) < MULTI_CONTEXT_MAX_REQUESTS
}

func (c *MultiClient) MultiCtxStreamingRequest(TextReader chan string, AlignmentResponseChannel chan StreamingOutputMultiCtxResponse, AudioResponsePipe io.Writer, voiceID string, modelID string, queries ...QueryFunc) error {
	driverActive := true // Driver shut down?
	driverError := false // Unexpected errors

	// // Eval context and capacity
	// if !c.hasMultiCtx(req.ContextID) {
	// 	if c.activeContextCount() >= MULTI_CONTEXT_MAX_REQUESTS {
	// 		return fmt.Errorf("active requests reached: %d", MULTI_CONTEXT_MAX_REQUESTS)
	// 	}
	// 	c.addMultiCtx(req.ContextID)
	// }

	// Make request
	url := fmt.Sprintf("%s/text-to-speech/%s/multi-stream-input?model_id=%s&inactivity_timeout=180&sync_alignment=true", ELEVEN_BASEURL_WSS, voiceID, modelID)
	headers := http.Header{}
	headers.Add("Accept", "*/*")
	headers.Add("Content-Type", JSON_CONTENT_TYPE)
	if c.apiKey != "" {
		headers.Add("xi-api-key", c.apiKey)
	}

	u, err := neturl.Parse(url)
	if err != nil {
		return err
	}

	q := u.Query()
	for _, qf := range queries {
		qf(&q)
	}
	u.RawQuery = q.Encode()

	conn, _, err := websocket.DefaultDialer.DialContext(c.ctx, u.String(), headers)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send initialization request and close initialization context
	var initReq TextToSpeechInputMultiStreamingRequest
	initCtx := cuid2.Generate()
	initReq = TextToSpeechInputMultiStreamingRequest{
		Text:      " ",
		ContextID: initCtx,
	}
	if err := conn.WriteJSON(initReq); err != nil {
		return err
	}
	initReq = TextToSpeechInputMultiStreamingRequest{
		CloseContext: true,
		ContextID:    initCtx,
	}
	if err := conn.WriteJSON(initReq); err != nil {
		return err
	}

	// Input watcher
	inputCtx, inputCancel := context.WithCancel(context.Background())

	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	// Response watching
	go func(wg *sync.WaitGroup, errCh chan<- error) {
		defer wg.Done()
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				if !driverActive {
					return
				}
				var input StreamingInputResponse
				var response StreamingOutputMultiCtxResponse
				if err := conn.ReadJSON(&input); err != nil {
					if driverActive {
						errCh <- err
						driverError = true
						inputCancel()
					}
					return
				}

				b, err := base64.StdEncoding.DecodeString(input.Audio)
				if err != nil {
					if driverActive {
						errCh <- err
						driverError = true
						inputCancel()
					}
					return
				}
				// Send audio through the pipeline
				if _, err := AudioResponsePipe.Write(b); err != nil {
					break
				}

				// Send non-audio via the response channel
				response = StreamingOutputMultiCtxResponse{
					IsFinal:             input.IsFinal,
					NormalizedAlignment: input.NormalizedAlignment,
					Alignment:           input.Alignment,
				}
				AlignmentResponseChannel <- response
			}
		}
	}(&wg, errCh)

	// Input watching
InputWatcher:
	for {
		select {
		case <-inputCtx.Done():
			driverActive = false
			break InputWatcher
		case <-c.ctx.Done():
			driverActive = false
			break InputWatcher
		case chunk, ok := <-TextReader:
			if !ok || !driverActive {
				break InputWatcher
			}
			ch := &TextToSpeechInputStreamingRequest{Text: chunk, ContextID: ""}
			if err := conn.WriteJSON(ch); err != nil {
				errCh <- err
				break InputWatcher
			}
		}
	}

	// Send final "" to close out TTS buffer
	if driverActive && !driverError {
		if err := conn.WriteJSON(map[string]string{"text": ""}); err != nil {
			if c.ctx.Err() == nil {
				errCh <- err
			}
		}
	}
	conn.Close()
	wg.Wait()

	// Errors?
	select {
	case readErr := <-errCh:
		//c.removeMultiCtx(req.ContextID)
		if driverActive || driverError {
			// Only send if the driver is active or the unexpected error flag is active
			return readErr
		} else {
			return nil
		}
	default:
	}
	//c.removeMultiCtx(req.ContextID)

	return nil
}

func LanguageCode(value string) QueryFunc {
	return func(q *neturl.Values) {
		q.Add("language_code", value)
	}
}

func OutputFormat(value string) QueryFunc {
	return func(q *neturl.Values) {
		q.Add("output_format", value)
	}
}

func SyncAlignment(value string) QueryFunc {
	return func(q *neturl.Values) {
		q.Add("sync_alignment", value)
	}
}

func InactivityTimeout(value string) QueryFunc {
	return func(q *neturl.Values) {
		q.Add("inactivity_timeout", value)
	}
}

func EnableSsmlParsing(value string) QueryFunc {
	return func(q *neturl.Values) {
		q.Add("enable_ssml_parsing", value)
	}
}
