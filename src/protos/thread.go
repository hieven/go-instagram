package protos

type Thread struct {
	ThreadID   string        `json:"thread_id"`
	ThreadV2ID string        `json:"thread_v2_id"`
	Users      []*ThreadUser `json:"users"`
	Items      []*ThreadItem `json:"items"`
	HasNewer   bool          `json:"has_newer"`
}
