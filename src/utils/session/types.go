package session

import "net/http"

type SessionManager interface {
	GetCookies() []*http.Cookie
	SetCookies([]*http.Cookie) error
}
