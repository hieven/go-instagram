package models

type Caption struct {
	Status       string `json:"status"`
	UserID       int    `json:"user_id"`
	CreatedAtUTC int64  `json:"created_at_utc"`
	CreatedAt    int64  `json:"created_at"`
	BitFlags     int    `json:"bit_flags"`
	User         User   `json:"user"`
	ContentType  string `json:"content_type"`
	Text         string `json:"text"`
	MediaID      int64  `json:"media_id"`
	Pk           int64  `json:"pk"`
	Type         int    `json:"type"`
}
