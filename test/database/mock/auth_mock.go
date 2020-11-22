package mock

import (
	"context"
	"ephelsa/my-career/pkg/auth/domain"
	"github.com/stretchr/testify/mock"
)

type authRepositoryMock struct {
	mock.Mock
}

func NewAuthRepositoryMock() *authRepositoryMock {
	return new(authRepositoryMock)
}

func (a *authRepositoryMock) IsUserRegistered(c context.Context, email string) (res bool, err error) {
	ret := a.Called(c, email)

	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		res = rf(c, email)
	} else {
		if ret.Get(0) != nil {
			res = ret.Get(0).(bool)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		err = rf(c, email)
	} else {
		err = ret.Error(1)
	}

	return
}

func (a *authRepositoryMock) IsUserRegistryConfirmed(c context.Context, email string) (res bool, err error) {
	ret := a.Called(c, email)

	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		res = rf(c, email)
	} else {
		if ret.Get(0) != nil {
			res = ret.Get(0).(bool)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		err = rf(c, email)
	} else {
		err = ret.Error(1)
	}

	return
}

func (a *authRepositoryMock) Register(c context.Context, r domain.Register) (res domain.RegisterSuccess, err error) {
	ret := a.Called(c, r)

	if rf, ok := ret.Get(0).(func(context.Context, domain.Register) domain.RegisterSuccess); ok {
		res = rf(c, r)
	} else {
		if ret.Get(0) != nil {
			res = ret.Get(0).(domain.RegisterSuccess)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Register) error); ok {
		err = rf(c, r)
	} else {
		err = ret.Error(1)
	}

	return
}

func (a *authRepositoryMock) IsAuthSuccess(c context.Context, auth domain.AuthCredentials) (res bool, err error) {
	ret := a.Called(c, auth)

	if rf, ok := ret.Get(0).(func(context.Context, domain.AuthCredentials) bool); ok {
		res = rf(c, auth)
	} else {
		if ret.Get(0) != nil {
			res = ret.Get(0).(bool)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.AuthCredentials) error); ok {
		err = rf(c, auth)
	} else {
		err = ret.Error(1)
	}

	return
}
