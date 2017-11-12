package protos

type LoginRequest struct {
	SignedBody      string `json:"signed_body"`
	IgSigKeyVersion string `json:"ig_sig_key_version"`
}

type InboxFeedRequest struct {
	Cursor string
}

type ThreadBroadcastTextRequest struct {
	UUID          string `json:"_uuid"`
	ThreadIDs     string `json:"thread_ids"`
	ClientContext string `json:"client_context"`
	Text          string `json:"text"`
}
