package protos

type LoginRequest struct {
	SignedBody      string `json:"signed_body"`
	IgSigKeyVersion string `json:"ig_sig_key_version"`
}

type TimelineFeedRequest struct {
	UserID    int64  // NOTE: required
	MaxID     string // NOTE: optional
	RankToken string // NOTE: optional
}

type InboxFeedRequest struct {
	Cursor string // NOTE: optional
}

type ThreadBroadcastTextRequest struct {
	UUID          string `json:"_uuid"`          // NOTE: optional
	ClientContext string `json:"client_context"` // NOTE: optional
	ThreadIDs     string `json:"thread_ids"`     // NOTE: required
	Text          string `json:"text"`           // NOTE: required
}

type ThreadBroadcastLinkRequest struct {
	UUID          string `json:"_uuid"`          // NOTE: optional
	ClientContext string `json:"client_context"` // NOTE: optional
	ThreadIDs     string `json:"thread_ids"`     // NOTE: required
	LinkText      string `json:"link_text"`      // NOTE: required
	LinkURLs      string `json:"link_urls"`      // NOTE: forbidden
}

type ThreadApproveAllRequest struct {
	UUID string `json:"_uuid"` // NOTE: optional
}
