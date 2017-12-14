package request

import (
	"github.com/stretchr/testify/mock"

	"github.com/parnurzeal/gorequest"
)

type MockCommon struct {
	mock.Mock
}

func (_m *MockCommon) WithDefaultHeader(rm *requestManager, req *gorequest.SuperAgent) *gorequest.SuperAgent {
	ret := _m.Called(rm, req)

	var r0 *gorequest.SuperAgent
	if rf, ok := ret.Get(0).(func(rm *requestManager, req *gorequest.SuperAgent) *gorequest.SuperAgent); ok {
		r0 = rf(rm, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorequest.SuperAgent)
		}
	}

	return r0
}
