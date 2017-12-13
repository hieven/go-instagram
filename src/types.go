package instagram

import (
	"context"

	"github.com/hieven/go-instagram/src/protos"
)

const (
	instaMsgLoginRequired = "login_required"
	instaStatusFail       = "fail"
)

type Instagram interface {
	Login(context.Context) error
	RememberMe(context.Context) error

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
	BroadcastShare(context.Context, *ThreadBroadcastShareRequest) (*protos.ThreadBroadcastShareResponse, error)
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
	UserID int64
	MaxID  string // NOTE: optional
}

type InboxFeedRequest struct {
	Cursor string // NOTE: optional
}

type ThreadApproveAllRequest struct{}

type ThreadBroadcastTextRequest struct {
	ThreadIDs string
	Text      string
}

type ThreadBroadcastLinkRequest struct {
	ThreadIDs string
	LinkText  string
}

type ThreadBroadcastShareRequest struct {
	ThreadIDs string
	MediaID   string
	Text      string
}

type ThreadShowRequest struct {
	ThreadID string
}

type MediaLikeRequest struct {
	MediaID string
}

type MediaUnlikeRequest struct {
	MediaID string
}

type LocationFeedRequest struct {
	Pk int64
}
