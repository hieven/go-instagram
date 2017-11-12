package instagram

import "github.com/hieven/go-instagram/src/protos"

type Instagram interface {
	Login() error
	Inbox() Inbox
	Timeline() Timeline
}

type Inbox interface {
	Feed() (*protos.InboxFeedResponse, error)
}

type Timeline interface {
	Feed() (string, error)
}
