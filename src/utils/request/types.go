package request

import "net/http"

type RequestManger interface {
	Get(url string) (resp *http.Response, body string, err error)
	Post(url string, data interface{}) (resp *http.Response, body string, err error)
}
