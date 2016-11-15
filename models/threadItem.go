package models

type ThreadItem struct {
	ID        string `json:"item_id"`
	UserID    int    `json:"user_id"`
	ItemType  string `json:"item_type"`
	Timestamp int    `json:"timestamp"`

	// depends on ItemType
	Placeholder placeholderSchema `json:"placeholder"`
	Text        string            `json:"text"`
	MediaShare  Media             `json:"media_share"`
	Location    Location          `json:"location"`
}

type placeholderSchema struct {
	IsLinked bool   `json:"is_linked"`
	Message  string `json:"message"`
	Title    string `json:"title"`
}
