package protos

type thread struct {
	ThreadID string        `json:"thread_id"`
	Users    []*ThreadUser `json:"users"`
	Items    []*ThreadItem `json:"items"`
	HasNewer bool          `json:"has_newer"`
}
