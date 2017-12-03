package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hieven/go-instagram/src/utils/session"
	"github.com/parnurzeal/gorequest"
)

type Common struct {
	mock.Mock
}

func (_m *Common) WithDefaultHeader(sessionManager session.SessionManager, req *gorequest.SuperAgent) *gorequest.SuperAgent {
	ret := _m.Called(sessionManager, req)

	var r0 *gorequest.SuperAgent
	if rf, ok := ret.Get(0).(func(sessionManager session.SessionManager, req *gorequest.SuperAgent) *gorequest.SuperAgent); ok {
		r0 = rf(sessionManager, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorequest.SuperAgent)
		}
	}

	return r0
}
