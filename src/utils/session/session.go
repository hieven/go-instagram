package session

import (
	"net/url"

	"github.com/hieven/go-instagram/src/config"
)

const (
	SchemeRedis = "redis"
)

func NewSession(cnf *config.Config) (SessionManager, error) {
	u, _ := url.Parse(cnf.SessionStorage)

	switch u.Scheme {
	case SchemeRedis:
		return newRedisSession(cnf)
	default:
		return newMemorySession()
	}
}
