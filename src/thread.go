package instagram

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	"github.com/hieven/go-instagram/src/utils/auth"
	"github.com/hieven/go-instagram/src/utils/request"
	"github.com/hieven/go-instagram/src/utils/text"
)

type thread struct {
	authManager    auth.AuthManager
	requestManager request.RequestManger
	textManager    text.TextManager
}

func (thread *thread) ApproveAll(ctx context.Context, req *ThreadApproveAllRequest) (*protos.ThreadApproveAllResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(constants.ThreadApproveAllEndpoint) // TODO: handle error

	internalReq := &protos.ThreadApproveAllRequest{
		UUID: thread.authManager.GenerateUUID(),
	}

	_, body, _ := thread.requestManager.Post(ctx, urlStru.String(), internalReq) // TODO: handle error

	result := &protos.ThreadApproveAllResponse{}
	json.Unmarshal([]byte(body), result)

	if result.Message == instaMsgLoginRequired {
		return result, ErrLoginRequired
	}

	if result.Status == instaStatusFail {
		return result, ErrUnknown
	}

	return result, nil
}

func (thread *thread) BroadcastText(ctx context.Context, req *ThreadBroadcastTextRequest) (*protos.ThreadBroadcastTextResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(constants.ThreadBroadcastTextEndpoint) // TODO: handle error

	internalReq := &protos.ThreadBroadcastTextRequest{
		ThreadIDs:     "[" + req.ThreadIDs + "]",
		Text:          req.Text,
		UUID:          thread.authManager.GenerateUUID(),
		ClientContext: thread.authManager.GenerateUUID(),
	}

	_, body, _ := thread.requestManager.Post(ctx, urlStru.String(), internalReq) // TODO: handle error

	result := &protos.ThreadBroadcastTextResponse{}
	json.Unmarshal([]byte(body), result)

	if result.Message == instaMsgLoginRequired {
		return result, ErrLoginRequired
	}

	if result.Status == instaStatusFail {
		return result, ErrUnknown
	}

	return result, nil
}

func (thread *thread) BroadcastLink(ctx context.Context, req *ThreadBroadcastLinkRequest) (*protos.ThreadBroadcastLinkResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(constants.ThreadBroadcastLinkEndpoint) // TODO: handle error

	internalReq := &protos.ThreadBroadcastLinkRequest{
		ThreadIDs:     "[" + req.ThreadIDs + "]",
		LinkText:      req.LinkText,
		LinkURLs:      `["` + thread.textManager.ExtractURL(req.LinkText) + `"]`,
		ClientContext: thread.authManager.GenerateUUID(),
		UUID:          thread.authManager.GenerateUUID(),
	}

	_, body, _ := thread.requestManager.Post(ctx, urlStru.String(), internalReq) // TODO: handle error

	result := &protos.ThreadBroadcastLinkResponse{}
	json.Unmarshal([]byte(body), result)

	if result.Message == instaMsgLoginRequired {
		return result, ErrLoginRequired
	}

	if result.Status == instaStatusFail {
		return result, ErrUnknown
	}

	return result, nil
}

func (thread *thread) BroadcastShare(ctx context.Context, req *ThreadBroadcastShareRequest) (*protos.ThreadBroadcastShareResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(constants.ThreadBroadcastShareEndpoint) // TODO: handle error

	internalReq := &protos.ThreadBroadcastShareRequest{
		ThreadIDs:     "[" + req.ThreadIDs + "]",
		Text:          req.Text,
		MediaID:       req.MediaID,
		ClientContext: thread.authManager.GenerateUUID(),
	}

	_, body, _ := thread.requestManager.Post(ctx, urlStru.String(), internalReq) // TODO: handle error

	result := &protos.ThreadBroadcastShareResponse{}
	json.Unmarshal([]byte(body), result)

	if result.Message == instaMsgLoginRequired {
		return result, ErrLoginRequired
	}

	if result.Status == instaStatusFail {
		return result, ErrUnknown
	}

	return result, nil
}

func (thread *thread) Show(ctx context.Context, req *ThreadShowRequest) (*protos.ThreadShowResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(fmt.Sprintf(constants.ThreadShowEndpoint, req.ThreadID)) // TODO: handle error

	_, body, _ := thread.requestManager.Get(ctx, urlStru.String()) // TODO: handle error

	result := &protos.ThreadShowResponse{}
	json.Unmarshal([]byte(body), result)

	if result.Message == instaMsgLoginRequired {
		return result, ErrLoginRequired
	}

	if result.Status == instaStatusFail {
		return result, ErrUnknown
	}

	if result.Thread == nil {
		return result, ErrUnknown
	}

	return result, nil
}
