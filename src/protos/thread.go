package protos

type thread struct {
	ThreadID string        `json:"thread_id"`
	Users    []*threadUser `json:"users"`
	Items    []*threadItem `json:"items"`
	HasNewer bool          `json:"has_newer"`
}
