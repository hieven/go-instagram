package request

import (
	"context"
	"net/http"

	"github.com/hieven/go-instagram/src/utils/session"
	"github.com/parnurzeal/gorequest"
)

type requestManager struct {
	cookies []*http.Cookie

	// utils
	sessionManager session.SessionManager
}

func New(sessionManager session.SessionManager) (RequestManger, error) {
	req := &requestManager{
		sessionManager: sessionManager,
	}
	return req, nil
}

func (request *requestManager) Get(ctx context.Context, url string) (*http.Response, string, error) {
	req := gorequest.New().
		Get(url)

	withDefaultHeader(request.sessionManager, req)

	resp, body, errs := req.End()

	var err error
	if len(errs) > 0 {
		err = errs[0]
	}

	return resp, body, err
}

func (request *requestManager) Post(ctx context.Context, url string, data interface{}) (*http.Response, string, error) {
	req := gorequest.New().
		Post(url).
		Type("multipart").
		SendStruct(data)

	withDefaultHeader(request.sessionManager, req)

	resp, body, errs := req.End()

	var err error
	if len(errs) > 0 {
		err = errs[0]
	}

	return resp, body, err
}
