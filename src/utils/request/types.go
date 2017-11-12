package request

import (
	"context"
	"net/http"
)

type RequestManger interface {
	Get(ctx context.Context, url string) (resp *http.Response, body string, err error)
	Post(ctx context.Context, url string, data interface{}) (resp *http.Response, body string, err error)
}
