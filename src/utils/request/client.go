package request

import (
	"net/http"

	"github.com/hieven/go-instagram/src/constants"
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

func (request *requestManager) SetCookies(cookies []*http.Cookie) {
	request.cookies = cookies
}

func (request *requestManager) Get(url string) (*http.Response, string, error) {
	req := gorequest.New().
		Get(url)

	request.withDefaultHeader(req)

	resp, body, errs := req.End()

	var err error
	if len(errs) > 0 {
		err = errs[0]
	}

	return resp, body, err
}

func (request *requestManager) Post(url string, data interface{}) (*http.Response, string, error) {
	req := gorequest.New().
		Post(url).
		Type("multipart").
		SendStruct(data)

	request.withDefaultHeader(req)

	resp, body, errs := req.End()

	var err error
	if len(errs) > 0 {
		err = errs[0]
	}

	return resp, body, err
}

func (request *requestManager) withDefaultHeader(req *gorequest.SuperAgent) *gorequest.SuperAgent {
	cookies := request.sessionManager.GetCookies()

	return req.
		Set("Connection", "close").
		Set("Accept", "*/*").
		Set("X-IG-Connection-Type", "WIFI").
		Set("X-IG-Capabilities", "3QI=").
		Set("Accept-Language", "en-US").
		Set("Host", constants.Hostname).
		Set("User-Agent", "Instagram "+constants.AppVersion+" Android (21/5.1.1; 401dpi; 1080x1920; Oppo; A31u; A31u; en_US)").
		AddCookies(cookies)
}
