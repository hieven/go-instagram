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
	Location() Location
}

type Timeline interface {
	Feed(context.Context, *TimelineFeedRequest) (*protos.TimelineFeedResponse, error)
}

type Inbox interface {
	Feed(context.Context, *InboxFeedRequest) (*protos.InboxFeedResponse, error)
}

type Thread interface {
	ApproveAll(context.Context, *ThreadApproveAllRequest) (*protos.ThreadApproveAllResponse, error)
	BroadcastText(context.Context, *ThreadBroadcastTextRequest) (*protos.ThreadBroadcastTextResponse, error)
	BroadcastLink(context.Context, *ThreadBroadcastLinkRequest) (*protos.ThreadBroadcastLinkResponse, error)
	Show(context.Context, *ThreadShowRequest) (*protos.ThreadShowResponse, error)
}

type Media interface {
	Like(context.Context, *MediaLikeRequest) (*protos.MediaLikeResponse, error)
	Unlike(context.Context, *MediaUnlikeRequest) (*protos.MediaUnlikeResponse, error)
}

type Location interface {
	Feed(context.Context, *LocationFeedRequest) (*protos.LocationFeedResponse, error)
}

type TimelineFeedRequest struct {
	UserID int64  // NOTE: required
	MaxID  string // NOTE: optional
}

type InboxFeedRequest struct {
	Cursor string // NOTE: optional
}

type ThreadApproveAllRequest struct{}

type ThreadBroadcastTextRequest struct {
	ThreadIDs string // NOTE: required
	Text      string // NOTE: required
}

type ThreadBroadcastLinkRequest struct {
	ThreadIDs string // NOTE: required
	LinkText  string // NOTE: required
}

type ThreadShowRequest struct {
	ThreadID string // NOTE: required
}

type MediaLikeRequest struct {
	MediaID string // NOTE: required
}

type MediaUnlikeRequest struct {
	MediaID string // NOTE: required
}

type LocationFeedRequest struct {
	Pk int64
}
