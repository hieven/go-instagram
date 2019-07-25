package instagram

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/hieven/go-instagram/src/utils/request"

	"github.com/hieven/go-instagram/src/utils/auth"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
)

type timeline struct {
	authManager    auth.AuthManager
	requestManager request.RequestManger
}

func (timeline *timeline) Feed(ctx context.Context, req *TimelineFeedRequest) (*protos.TimelineFeedResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(constants.TimelineFeedEndpoint) // TODO: handle error

	internalReq := &protos.TimelineFeedRequest{
		UserID:    req.UserID,
		MaxID:     req.MaxID,
		RankToken: timeline.authManager.GenerateRankToken(req.UserID),
	}

	_, body, _ := timeline.requestManager.Post(ctx, urlStru.String(), internalReq) // TODO: handle error
	result := &protos.TimelineFeedResponse{}
	json.Unmarshal([]byte(body), result) // TODO: handle error
	return result, nil
}
