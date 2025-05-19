package elevenlabs

type UserAndCapacity struct {
	UserID       string       `json:"user_id"`
	Subscription Subscription `json:"subscription"`
	HasCapacity  bool         `json:"has_capacity"`
}
