package mock

import (
	"context"
	authDomain "ephelsa/my-career/pkg/auth/domain"
	"ephelsa/my-career/pkg/user/domain"
	"github.com/stretchr/testify/mock"
)

type userMockLocalRepo struct {
	mock.Mock
}

func NewUserLocalRepositoryMock() *userMockLocalRepo {
	return new(userMockLocalRepo)
}

func (u *userMockLocalRepo) InformationByEmail(ctx context.Context, email string) (res domain.User, err error) {
	ret := u.Called(ctx, email)

	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		res = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			res = ret.Get(0).(domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		err = rf(ctx, email)
	} else {
		err = ret.Error(1)
	}

	return
}

func (u *userMockLocalRepo) StoreUserInformation(ctx context.Context, user domain.User) (res authDomain.RegisterSuccess, err error) {
	ret := u.Called(ctx, user)

	if rf, ok := ret.Get(0).(func(context.Context, domain.User) authDomain.RegisterSuccess); ok {
		res = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			res = ret.Get(0).(authDomain.RegisterSuccess)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		err = rf(ctx, user)
	} else {
		err = ret.Error(1)
	}

	return
}
