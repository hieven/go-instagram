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
	Media() Media
}

type Timeline interface {
	Feed(context.Context, *protos.TimelineFeedRequest) (*protos.TimelineFeedResponse, error)
}

type Inbox interface {
	Feed(context.Context, *protos.InboxFeedRequest) (*protos.InboxFeedResponse, error)
}

type Thread interface {
	ApproveAll(context.Context, *protos.ThreadApproveAllRequest) (*protos.ThreadApproveAllResponse, error)
	BroadcastText(context.Context, *protos.ThreadBroadcastTextRequest) (*protos.ThreadBroadcastTextResponse, error)
	BroadcastLink(context.Context, *protos.ThreadBroadcastLinkRequest) (*protos.ThreadBroadcastLinkResponse, error)
	Show(context.Context, *protos.ThreadShowRequest) (*protos.ThreadShowResponse, error)
}

type Media interface {
	Like(context.Context, *protos.MediaLikeRequest) (*protos.MediaLikeResponse, error)
	Unlike(context.Context, *protos.MediaUnlikeRequest) (*protos.MediaUnlikeResponse, error)
}
