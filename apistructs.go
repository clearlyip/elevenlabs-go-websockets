package elevenlabs

type UserData struct {
	UserID                         string             `json:"user_id"`
	Subscription                   Subscription       `json:"subscription"`
	IsNewUser                      bool               `json:"is_new_user"`
	CanUseDelayedPaymentMethods    bool               `json:"can_use_delayed_payment_methods"`
	IsOnboardingCompleted          bool               `json:"is_onboarding_completed"`
	IsOnboardingChecklistCompleted bool               `json:"is_onboarding_checklist_completed"`
	SubscriptionExtras             SubscriptionExtras `json:"subscription_extras"`
	XIApiKey                       string             `json:"xi_api_key"`
	FirstName                      string             `json:"first_name"`
	IsAPIKeyHashed                 bool               `json:"is_api_key_hashed"`
	XIApiKeyPreview                string             `json:"xi_api_key_preview"`
	ReferralLinkCode               string             `json:"referral_link_code"`
	PartnerstackPartnerDefaultLink string             `json:"partnerstack_partner_default_link"`
}

type Subscription struct {
	Tier                           string `json:"tier"`
	CharacterCount                 int    `json:"character_count"`
	CharacterLimit                 int    `json:"character_limit"`
	CanExtendCharacterLimit        bool   `json:"can_extend_character_limit"`
	AllowedToExtendCharacterLimit  bool   `json:"allowed_to_extend_character_limit"`
	VoiceSlotsUsed                 int    `json:"voice_slots_used"`
	ProfessionalVoiceSlotsUsed     int    `json:"professional_voice_slots_used"`
	VoiceLimit                     int    `json:"voice_limit"`
	VoiceAddEditCounter            int    `json:"voice_add_edit_counter"`
	ProfessionalVoiceLimit         int    `json:"professional_voice_limit"`
	CanExtendVoiceLimit            bool   `json:"can_extend_voice_limit"`
	CanUseInstantVoiceCloning      bool   `json:"can_use_instant_voice_cloning"`
	CanUseProfessionalVoiceCloning bool   `json:"can_use_professional_voice_cloning"`
	Status                         string `json:"status"`
	MaxCharacterLimitExtension     int    `json:"max_character_limit_extension"`
	NextCharacterCountResetUnix    int64  `json:"next_character_count_reset_unix"`
	MaxVoiceAddEdits               int    `json:"max_voice_add_edits"`
	Currency                       string `json:"currency"`
	BillingPeriod                  string `json:"billing_period"`
	CharacterRefreshPeriod         string `json:"character_refresh_period"`
}

type SubscriptionExtras struct {
	Concurrency                                    int        `json:"concurrency"`
	ConvaiConcurrency                              int        `json:"convai_concurrency"`
	ForceLoggingDisabled                           bool       `json:"force_logging_disabled"`
	CanRequestManualProVoiceVerification           bool       `json:"can_request_manual_pro_voice_verification"`
	CanBypassVoiceCaptcha                          bool       `json:"can_bypass_voice_captcha"`
	Moderation                                     Moderation `json:"moderation"`
	ConvaiCharsPerMinute                           int        `json:"convai_chars_per_minute"`
	ConvaiASRCharsPerMinute                        int        `json:"convai_asr_chars_per_minute"`
	UnusedCharactersRolledOverFromPreviousPeriod   int        `json:"unused_characters_rolled_over_from_previous_period"`
	OverusedCharactersRolledOverFromPreviousPeriod int        `json:"overused_characters_rolled_over_from_previous_period"`
	Usage                                          Usage      `json:"usage"`
}

type Moderation struct {
	IsInProbation                         bool   `json:"is_in_probation"`
	EnterpriseCheckNogoVoice              bool   `json:"enterprise_check_nogo_voice"`
	EnterpriseCheckBlockNogoVoice         bool   `json:"enterprise_check_block_nogo_voice"`
	NeverLiveModerate                     bool   `json:"never_live_moderate"`
	NogoVoiceSimilarVoiceUploadCount      int    `json:"nogo_voice_similar_voice_upload_count"`
	EnterpriseBackgroundModerationEnabled bool   `json:"enterprise_background_moderation_enabled"`
	OnWatchlist                           bool   `json:"on_watchlist"`
	SafetyStatus                          string `json:"safety_status"`
	WarningStatus                         string `json:"warning_status"`
}

type Usage struct {
	RolloverCreditsQuota          int `json:"rollover_credits_quota"`
	SubscriptionCycleCreditsQuota int `json:"subscription_cycle_credits_quota"`
	ManuallyGiftedCreditsQuota    int `json:"manually_gifted_credits_quota"`
	RolloverCreditsUsed           int `json:"rollover_credits_used"`
	SubscriptionCycleCreditsUsed  int `json:"subscription_cycle_credits_used"`
	ManuallyGiftedCreditsUsed     int `json:"manually_gifted_credits_used"`
	PaidUsageBasedCreditsUsed     int `json:"paid_usage_based_credits_used"`
	ActualReportedCredits         int `json:"actual_reported_credits"`
}

type StreamingAlignmentSegment struct {
	CharStartTimesMs []int    `json:"charStartTimesMs"`
	CharDurationsMs  []int    `json:"charDurationsMs"`
	Chars            []string `json:"chars"`
}

type ListVoicesParams struct {
	PageSize            int                `url:"page_size,omitempty"` // <=100, default 30
	Category            VoiceParamCategory `url:"category,omitempty"`  // "professional", etc.
	Gender              string             `url:"gender,omitempty"`
	Age                 string             `url:"age,omitempty"`
	Accent              string             `url:"accent,omitempty"`
	Language            string             `url:"language,omitempty"`
	Locale              string             `url:"locale,omitempty"`
	Search              string             `url:"search,omitempty"`
	UseCases            string             `url:"use_cases,omitempty"`
	Descriptives        string             `url:"descriptives,omitempty"`
	Featured            bool               `url:"featured,omitempty"` // default false
	MinNoticePeriodDays int                `url:"min_notice_period_days,omitempty"`
	IncludeCustomRates  bool               `url:"include_custom_rates,omitempty"`
	ReaderAppEnabled    bool               `url:"reader_app_enabled,omitempty"` // default false
	OwnerID             string             `url:"owner_id,omitempty"`
	Sort                string             `url:"sort,omitempty"`
	Page                int                `url:"page,omitempty"` // default 0
}
type VoiceParamCategory string

type ListVoicesResponse struct {
	Voices     []Voice `json:"voices"`
	HasMore    bool    `json:"has_more"`
	LastSortID string  `json:"last_sort_id"`
}

type Voice struct {
	PublicOwnerID                string                  `json:"public_owner_id"`
	VoiceID                      string                  `json:"voice_id"`
	DateUnix                     int64                   `json:"date_unix"`
	Name                         string                  `json:"name"`
	Accent                       string                  `json:"accent"`
	Gender                       string                  `json:"gender"`
	Age                          string                  `json:"age"`
	Descriptive                  string                  `json:"descriptive"`
	UseCase                      string                  `json:"use_case"`
	Category                     string                  `json:"category"`
	UsageCharacterCount1Y        int64                   `json:"usage_character_count_1y"`
	UsageCharacterCount7D        int64                   `json:"usage_character_count_7d"`
	PlayAPIUsageCharacterCount1Y int64                   `json:"play_api_usage_character_count_1y"`
	ClonedByCount                int                     `json:"cloned_by_count"`
	FreeUsersAllowed             bool                    `json:"free_users_allowed"`
	LiveModerationEnabled        bool                    `json:"live_moderation_enabled"`
	Featured                     bool                    `json:"featured"`
	Language                     string                  `json:"language"`
	Locale                       string                  `json:"locale"`
	Description                  string                  `json:"description"`
	PreviewURL                   string                  `json:"preview_url"`
	Rate                         float64                 `json:"rate"`
	FiatRate                     float64                 `json:"fiat_rate"`
	VerifiedLanguages            []VoiceVerifiedLanguage `json:"verified_languages"`
	NoticePeriod                 int                     `json:"notice_period"`
	InstagramUsername            string                  `json:"instagram_username"`
	TwitterUsername              string                  `json:"twitter_username"`
	YouTubeUsername              string                  `json:"youtube_username"`
	TikTokUsername               string                  `json:"tiktok_username"`
	ImageURL                     string                  `json:"image_url"`
	IsAddedByUser                bool                    `json:"is_added_by_user"`
}

type VoiceVerifiedLanguage struct {
	Language string `json:"language"`
	ModelID  string `json:"model_id"`
	Accent   string `json:"accent"`
}

type GetVoiceVoice struct {
	VoiceID                 string                     `json:"voice_id"`
	Name                    string                     `json:"name"`
	Samples                 []GetVoiceSample           `json:"samples"`
	Category                string                     `json:"category"`
	FineTuning              GetVoiceFineTuning         `json:"fine_tuning"`
	Labels                  GetVoiceLabels             `json:"labels"`
	Description             string                     `json:"description"`
	PreviewURL              string                     `json:"preview_url"`
	AvailableForTiers       []string                   `json:"available_for_tiers"`
	Settings                GetVoiceSettings           `json:"settings"`
	Sharing                 GetVoiceSharing            `json:"sharing"`
	HighQualityBaseModelIDs []string                   `json:"high_quality_base_model_ids"`
	VerifiedLanguages       []GetVoiceVerifiedLanguage `json:"verified_languages"`
	SafetyControl           *string                    `json:"safety_control"`
	VoiceVerification       GetVoiceVoiceVerification  `json:"voice_verification"`
	PermissionOnResource    *string                    `json:"permission_on_resource"`
	IsOwner                 *bool                      `json:"is_owner"`
	IsLegacy                bool                       `json:"is_legacy"`
	IsMixed                 bool                       `json:"is_mixed"`
	CreatedAtUnix           int64                      `json:"created_at_unix"`
}

type GetVoiceSample struct {
	SampleID                string                    `json:"sample_id"`
	FileName                string                    `json:"file_name"`
	MimeType                string                    `json:"mime_type"`
	SizeBytes               int64                     `json:"size_bytes"`
	Hash                    string                    `json:"hash"`
	DurationSecs            float64                   `json:"duration_secs"`
	RemoveBackgroundNoise   bool                      `json:"remove_background_noise"`
	HasIsolatedAudio        bool                      `json:"has_isolated_audio"`
	HasIsolatedAudioPreview bool                      `json:"has_isolated_audio_preview"`
	SpeakerSeparation       GetVoiceSpeakerSeparation `json:"speaker_separation"`
	TrimStart               int                       `json:"trim_start"`
	TrimEnd                 int                       `json:"trim_end"`
}

type GetVoiceSpeakerSeparation struct {
	VoiceID  string `json:"voice_id"`
	SampleID string `json:"sample_id"`
	Status   string `json:"status"`
}

type GetVoiceFineTuning struct {
	IsAllowedToFineTune                    bool                          `json:"is_allowed_to_fine_tune"`
	State                                  map[string]string             `json:"state"`
	VerificationFailures                   []string                      `json:"verification_failures"`
	VerificationAttemptsCount              int                           `json:"verification_attempts_count"`
	ManualVerificationRequested            bool                          `json:"manual_verification_requested"`
	Language                               string                        `json:"language"`
	Progress                               map[string]float64            `json:"progress"`
	Message                                map[string]string             `json:"message"`
	DatasetDurationSeconds                 *float64                      `json:"dataset_duration_seconds"`
	VerificationAttempts                   []GetVoiceVerificationAttempt `json:"verification_attempts"`
	SliceIDs                               []string                      `json:"slice_ids"`
	ManualVerification                     *GetVoiceManualVerification   `json:"manual_verification"`
	MaxVerificationAttempts                int                           `json:"max_verification_attempts"`
	NextMaxVerificationAttemptsResetUnixMs int64                         `json:"next_max_verification_attempts_reset_unix_ms"`
}

type GetVoiceVerificationAttempt struct {
	Text                string            `json:"text"`
	DateUnix            int64             `json:"date_unix"`
	Accepted            bool              `json:"accepted"`
	Similarity          float64           `json:"similarity"`
	LevenshteinDistance int               `json:"levenshtein_distance"`
	Recording           GetVoiceRecording `json:"recording"`
}

type GetVoiceRecording struct {
	RecordingID    string `json:"recording_id"`
	MimeType       string `json:"mime_type"`
	SizeBytes      int64  `json:"size_bytes"`
	UploadDateUnix int64  `json:"upload_date_unix"`
	Transcription  string `json:"transcription"`
}

type GetVoiceManualVerification struct {
	ExtraText       string                           `json:"extra_text"`
	RequestTimeUnix int64                            `json:"request_time_unix"`
	Files           []GetVoiceManualVerificationFile `json:"files"`
}

type GetVoiceManualVerificationFile struct {
	FileID         string `json:"file_id"`
	FileName       string `json:"file_name"`
	MimeType       string `json:"mime_type"`
	SizeBytes      int64  `json:"size_bytes"`
	UploadDateUnix int64  `json:"upload_date_unix"`
}

type GetVoiceLabels struct {
	Language    string `json:"language"`
	Descriptive string `json:"descriptive"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Accent      string `json:"accent"`
	UseCase     string `json:"use_case"`
	Locale      string `json:"locale"`
}

type GetVoiceSettings struct {
	Stability       float64 `json:"stability"`
	SimilarityBoost float64 `json:"similarity_boost"`
	Style           float64 `json:"style"`
	UseSpeakerBoost bool    `json:"use_speaker_boost"`
	Speed           float64 `json:"speed"`
}

type GetVoiceSharing struct {
	Status                  string                       `json:"status"`
	HistoryItemSampleID     *string                      `json:"history_item_sample_id"`
	DateUnix                int64                        `json:"date_unix"`
	WhitelistedEmails       []string                     `json:"whitelisted_emails"`
	PublicOwnerID           string                       `json:"public_owner_id"`
	OriginalVoiceID         string                       `json:"original_voice_id"`
	FinancialRewardsEnabled bool                         `json:"financial_rewards_enabled"`
	FreeUsersAllowed        bool                         `json:"free_users_allowed"`
	LiveModerationEnabled   bool                         `json:"live_moderation_enabled"`
	Rate                    float64                      `json:"rate"`
	FiatRate                *float64                     `json:"fiat_rate"`
	NoticePeriod            int                          `json:"notice_period"`
	DisableAtUnix           *int64                       `json:"disable_at_unix"`
	VoiceMixingAllowed      bool                         `json:"voice_mixing_allowed"`
	Featured                bool                         `json:"featured"`
	Category                string                       `json:"category"`
	ReaderAppEnabled        *bool                        `json:"reader_app_enabled"`
	ImageURL                string                       `json:"image_url"`
	BanReason               *string                      `json:"ban_reason"`
	LikedByCount            int                          `json:"liked_by_count"`
	ClonedByCount           int                          `json:"cloned_by_count"`
	Name                    string                       `json:"name"`
	Description             string                       `json:"description"`
	Labels                  GetVoiceSharingLabels        `json:"labels"`
	ReviewStatus            string                       `json:"review_status"`
	ReviewMessage           *string                      `json:"review_message"`
	EnabledInLibrary        bool                         `json:"enabled_in_library"`
	InstagramUsername       *string                      `json:"instagram_username"`
	TwitterUsername         *string                      `json:"twitter_username"`
	YouTubeUsername         *string                      `json:"youtube_username"`
	TikTokUsername          *string                      `json:"tiktok_username"`
	ModerationCheck         *GetVoiceModerationCheck     `json:"moderation_check"`
	ReaderRestrictedOn      []GetVoiceRestrictedResource `json:"reader_restricted_on"`
}

type GetVoiceSharingLabels struct {
	Language    string `json:"language"`
	Descriptive string `json:"descriptive"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Accent      string `json:"accent"`
	UseCase     string `json:"use_case"`
	Locale      string `json:"locale"`
}

type GetVoiceModerationCheck struct {
	DateCheckedUnix  int64     `json:"date_checked_unix"`
	NameValue        string    `json:"name_value"`
	NameCheck        bool      `json:"name_check"`
	DescriptionValue string    `json:"description_value"`
	DescriptionCheck bool      `json:"description_check"`
	SampleIDs        []string  `json:"sample_ids"`
	SampleChecks     []float64 `json:"sample_checks"`
	CaptchaIDs       []string  `json:"captcha_ids"`
	CaptchaChecks    []float64 `json:"captcha_checks"`
}

type GetVoiceRestrictedResource struct {
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id"`
}

type GetVoiceVerifiedLanguage struct {
	Language   string `json:"language"`
	ModelID    string `json:"model_id"`
	Accent     string `json:"accent"`
	Locale     string `json:"locale"`
	PreviewURL string `json:"preview_url"`
}

type GetVoiceVoiceVerification struct {
	RequiresVerification      bool                          `json:"requires_verification"`
	IsVerified                bool                          `json:"is_verified"`
	VerificationFailures      []string                      `json:"verification_failures"`
	VerificationAttemptsCount int                           `json:"verification_attempts_count"`
	Language                  *string                       `json:"language"`
	VerificationAttempts      []GetVoiceVerificationAttempt `json:"verification_attempts"`
}
type TextToSpeechInputMultiStreamingRequest struct {
	ContextID        string            `json:"context_id,omitempty"`
	CloseContext     bool              `json:"close_context,omitempty"`
	CloseSocket      bool              `json:"close_socket,omitempty"`
	Text             string            `json:"text,omitempty"`
	Flush            bool              `json:"flush,omitempty"`
	VoiceSettings    *VoiceSettings    `json:"voice_settings,omitempty"`
	GenerationConfig *GenerationConfig `json:"generation_config,omitempty"`
	//PronunciationDictionaryLocators PronunciationDictionaryLocators `json:"pronunciation_dictionary_locators,omitempty"`
}

type PronunciationDictionaryLocators struct {
	DictionaryId string `json:"dictionary_id"`
	VersionId    string `json:"version_id"`
}
