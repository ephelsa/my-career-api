package mock

import (
	"context"
	"ephelsa/my-career/pkg/institutiontype/domain"
	"github.com/stretchr/testify/mock"
)

type institutionTypeRepositoryMock struct {
	mock.Mock
}

func NewInstitutionTypeRepositoryMock() *institutionTypeRepositoryMock {
	return new(institutionTypeRepositoryMock)
}

func (i *institutionTypeRepositoryMock) FetchAll(c context.Context) (result []domain.InstitutionType, err error) {
	ret := i.Called(c)

	if rf, ok := ret.Get(0).(func(context.Context) []domain.InstitutionType); ok {
		result = rf(c)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).([]domain.InstitutionType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		err = rf(c)
	} else {
		err = ret.Error(1)
	}

	return
}
