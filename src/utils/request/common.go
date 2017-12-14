package request

import (
	"github.com/hieven/go-instagram/src/constants"
	"github.com/parnurzeal/gorequest"
)

var withDefaultHeader = func(rm *requestManager, req *gorequest.SuperAgent) *gorequest.SuperAgent {
	if len(rm.cookies) == 0 {
		rm.cookies = rm.sessionManager.GetCookies()
	}

	return req.
		Set("Connection", "close").
		Set("Accept", "*/*").
		Set("X-IG-Connection-Type", "WIFI").
		Set("X-IG-Capabilities", "3QI=").
		Set("Accept-Language", "en-US").
		Set("Host", constants.Hostname).
		Set("User-Agent", "Instagram "+constants.AppVersion+" Android (21/5.1.1; 401dpi; 1080x1920; Oppo; A31u; A31u; en_US)").
		AddCookies(rm.cookies)
}
