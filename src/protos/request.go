package protos

type LoginRequest struct {
	SignedBody      string `json:"signed_body"`
	IgSigKeyVersion string `json:"ig_sig_key_version"`
}

type InboxFeedRequest struct {
	Cursor string // NOTE: optional
}

type ThreadBroadcastTextRequest struct {
	UUID          string `json:"_uuid"`          // NOTE: optional
	ClientContext string `json:"client_context"` // NOTE: optional
	ThreadIDs     string `json:"thread_ids"`
	Text          string `json:"text"`
}

type ThreadBroadcastLinkRequest struct {
	UUID          string `json:"_uuid"`          // NOTE: optional
	ClientContext string `json:"client_context"` // NOTE: optional
	ThreadIDs     string `json:"thread_ids"`
	LinkText      string `json:"link_text"`
	LinkURLs      string `json:"link_urls"` // NOTE: forbidden
}
