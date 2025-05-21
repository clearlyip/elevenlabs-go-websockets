// Multi-context
package elevenlabs

import (
	"context"
	"io"
	"sync"
	"time"
)

type MultiClient struct {
	apiKey                   string
	timeout                  time.Duration
	ctx                      context.Context
	activeRequests           map[string]struct{} // Active multi-context requests
	cmu                      sync.RWMutex
	TextReader               chan string
	AlignmentResponseChannel chan StreamingOutputMultiCtxResponse
}

// Multi-Context Websocket Session
func NewMultiContextSession(ctx context.Context, apiKey string, reqTimeout time.Duration, TextReader chan string, AlignmentResponseChannel chan StreamingOutputMultiCtxResponse, AudioResponsePipe io.Writer, voiceID string, modelID string, req TextToSpeechInputMultiStreamingRequest, queries ...QueryFunc) *MultiClient {

	return &MultiClient{
		apiKey:                   apiKey,
		timeout:                  reqTimeout,
		ctx:                      ctx,
		activeRequests:           make(map[string]struct{}),
		TextReader:               TextReader,
		AlignmentResponseChannel: AlignmentResponseChannel,
	}

	// err := c.MultiCtxStreamingRequest(TextReader, AlignmentResponseChannel, AudioResponsePipe, voiceID, modelID, queries...)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
