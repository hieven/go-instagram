package session

import (
	"net/http"
)

type MemorySession struct {
	cookies []*http.Cookie
}

func newMemorySession() (*MemorySession, error) {
	session := MemorySession{}

	return &session, nil
}

func (session *MemorySession) GetCookies() []*http.Cookie {
	return session.cookies
}

func (session *MemorySession) SetCookies(cookies []*http.Cookie) error {
	session.cookies = cookies
	return nil
}
