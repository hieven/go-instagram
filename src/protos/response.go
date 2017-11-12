package protos

type LoginResponse struct {
	LoggedInUser *loggedInUser `json:"logged_in_user"`
	Status       string        `json:"status"`
	Message      string        `json:"message"`
}

type InboxFeedResponse struct {
	Inbox                *inbox `json:"inbox"`
	Status               string `json:"status"`
	PendingRequestsTotal int    `json:"pending_requests_total"`
	SeqID                int    `json:"seq_id"`
	Message              string `json:"message"`
	// PendingRequestsUsers []string `json:"pending_requests_users"`
}
