package protos

type threadItem struct {
	ID        string `json:"item_id"`
	UserID    int    `json:"user_id"`
	ItemType  string `json:"item_type"`
	Timestamp int    `json:"timestamp"`

	// NOTE: depends on ItemType
	Placeholder *placeholderSchema `json:"placeholder"`
	Text        string             `json:"text"`
	MediaShare  *media             `json:"media_share"`
	Location    *location          `json:"location"`
}

type placeholderSchema struct {
	IsLinked bool   `json:"is_linked"`
	Message  string `json:"message"`
	Title    string `json:"title"`
}
