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
