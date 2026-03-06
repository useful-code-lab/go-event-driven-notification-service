package domain

type Notification struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	Type string `json:"type"`
	Message string `json:"message"`
}