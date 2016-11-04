package utils

import (
	"github.com/hieven/go-instagram/constants"
	"github.com/parnurzeal/gorequest"
)

// type Request struct {
// 	*gorequest.SuperAgent
// }

// func New() (request *Request) {
// 	request = &Request{
// 		SuperAgent: gorequest.New(),
// 	}

// 	return request
// }

// func (request Request) Done() (resp gorequest.Response, body string, err []error) {
// 	resp, body, err = request.
// 		Set("X-IG-Connection-Type", "WIFI").
// 		Set("X-IG-Capabilities", "3QI=").
// 		Set("Accept-Language", "en-US").
// 		Set("Host", constants.HOSTNAME).
// 		Set("User-Agent", "Instagram "+constants.APP_VERSION+" Android (22/6.0.1; 515dpi; 2560x1440; Proscan; PLT1077G; PLT1077G; en_US)").
// 		End()

// 	return resp, body, err
// }

func WrapRequest(request *gorequest.SuperAgent) (resp gorequest.Response, body string, err []error) {
	resp, body, err = request.
		Set("Connection", "close").
		Set("Accept", "*/*").
		Set("X-IG-Connection-Type", "WIFI").
		Set("X-IG-Capabilities", "3QI=").
		Set("Accept-Language", "en-US").
		Set("Host", constants.HOSTNAME).
		Set("User-Agent", "Instagram "+constants.APP_VERSION+" Android (21/5.1.1; 401dpi; 1080x1920; Oppo; A31u; A31u; en_US)").
		End()

	return resp, body, err
}
