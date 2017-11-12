package instagram

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	"github.com/hieven/go-instagram/src/utils/auth"
	"github.com/hieven/go-instagram/src/utils/request"
)

type thread struct {
	authManager    auth.AuthManager
	requestManager request.RequestManger
}

func (thread *thread) BroadcastText(ctx context.Context, req *protos.ThreadBroadcastTextRequest) (*protos.ThreadBroadcastTextResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	if req.ClientContext == "" {
		req.ClientContext = thread.authManager.GenerateUUID()
	}
	if req.UUID == "" {
		req.UUID = thread.authManager.GenerateUUID()
	}
	req.ThreadIDs = "[" + req.ThreadIDs + "]"

	urlStru, _ := url.Parse(constants.ThreadBroadcastTextEndpoint) // TODO: handle error

	_, body, _ := thread.requestManager.Post(ctx, urlStru.String(), req) // TODO: handle error

	result := &protos.ThreadBroadcastTextResponse{}
	json.Unmarshal([]byte(body), result) // TODO: handle error

	return result, nil
}
