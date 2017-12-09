package instagram

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockCommon struct {
	mock.Mock
}

func (_m *MockCommon) Login(ig *instagram, ctx context.Context) error {
	ret := _m.Called(ig, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(ig *instagram, ctx context.Context) error); ok {
		r0 = rf(ig, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}

	return r0
}
