package instagram

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	"github.com/hieven/go-instagram/src/utils/request"
)

type inbox struct {
	requestManager request.RequestManger
}

func (inbox *inbox) Feed(ctx context.Context, req *InboxFeedRequest) (*protos.InboxFeedResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(constants.InboxEndpoint) // TODO: handle error
	query := urlStru.Query()

	if req.Cursor != "" {
		query.Set("cursor", req.Cursor)
	}

	urlStru.RawQuery = query.Encode()

	_, body, _ := inbox.requestManager.Get(ctx, urlStru.String()) // TODO: handle error

	result := &protos.InboxFeedResponse{}
	json.Unmarshal([]byte(body), result)

	if result.Message == instaMsgLoginRequired {
		return result, ErrLoginRequired
	}

	if result.Status == instaStatusFail {
		return result, ErrUnknown
	}

	if result.Inbox == nil {
		return result, ErrUnknown
	}

	return result, nil
}
