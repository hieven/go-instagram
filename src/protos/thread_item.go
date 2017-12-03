package protos

type ThreadItem struct {
	ID        string `json:"item_id"`
	UserID    int    `json:"user_id"`
	ItemType  string `json:"item_type"`
	Timestamp int    `json:"timestamp"`

	// NOTE: depends on ItemType
	Placeholder *PlaceholderSchema `json:"placeholder"`
	Text        string             `json:"text"`
	MediaShare  *Media             `json:"media_share"`
	Location    *Location          `json:"location"`
}

type PlaceholderSchema struct {
	IsLinked bool   `json:"is_linked"`
	Message  string `json:"message"`
	Title    string `json:"title"`
}
