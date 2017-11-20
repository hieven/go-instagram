package protos

type Inbox struct {
	Threads       []*ThreadItem `json:"threads"`
	HasOlder      bool          `json:"has_older"`
	OldestCursor  string        `json:"oldest_cursor"`
	UnseenCount   int           `json:"unseen_count"`
	UnseenCountTs int           `json:"unseen_count_ts"`
}
