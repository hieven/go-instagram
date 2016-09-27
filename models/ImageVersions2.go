package models

type ImageVersions2 struct {
	Candidates []Candidate `json:"candidates"`
}
type Candidate struct {
	URL    string
	Width  int
	Height int
}
