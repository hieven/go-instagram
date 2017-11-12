package instagram

import (
	"context"

	"github.com/hieven/go-instagram/src/protos"
)

type Instagram interface {
	Login(context.Context) error

	Timeline() Timeline
	Inbox() Inbox
	Thread() Thread
}

type Timeline interface {
	Feed() (string, error)
}

type Inbox interface {
	Feed(context.Context, *protos.InboxFeedRequest) (*protos.InboxFeedResponse, error)
}

type Thread interface {
	BroadcastText(context.Context, *protos.ThreadBroadcastTextRequest) (*protos.ThreadBroadcastTextResponse, error)
	BroadcastLink(context.Context, *protos.ThreadBroadcastLinkRequest) (*protos.ThreadBroadcastLinkResponse, error)
}
