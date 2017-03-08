package session

import (
	"net/http"
	"net/url"

	"github.com/hieven/go-instagram/constants"
)

type MemorySession struct {
	cookies []*http.Cookie
}

func NewMemorySession() (*MemorySession, error) {
	session := MemorySession{}

	return &session, nil
}

func (session *MemorySession) GetCookies() []*http.Cookie {
	return session.cookies
}

func (session *MemorySession) SetCookies(client *http.Client) error {
	u, err := url.Parse(constants.HOST)

	if err != nil {
		return err
	}

	session.cookies = client.Jar.Cookies(u)

	return nil
}
