package session

import (
	"net/http"
	"net/url"

	"github.com/hieven/go-instagram/config"
)

type Session interface {
	GetCookies() []*http.Cookie
	SetCookies(*http.Client) error
}

func NewSession(cnf *config.Config) (Session, error) {
	u, _ := url.Parse(cnf.SessionStorage)

	switch u.Scheme {
	case "redis":
		return NewRedisSession(cnf)
	default:
		return NewMemorySession()
	}
}
