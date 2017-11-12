package protos

type defaultResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LoginResponse struct {
	defaultResponse
	LoggedInUser *loggedInUser `json:"logged_in_user"`
}

type InboxFeedResponse struct {
	defaultResponse
	Inbox                *inbox `json:"inbox"`
	PendingRequestsTotal int    `json:"pending_requests_total"`
	SeqID                int    `json:"seq_id"`
	// PendingRequestsUsers []string `json:"pending_requests_users"`
}

type ThreadBroadcastTextResponse struct {
	defaultResponse
	Threads []*thread `json:"threads"`
}
