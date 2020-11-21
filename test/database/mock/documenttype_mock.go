package mock

import (
	"context"
	"ephelsa/my-career/pkg/documenttype/domain"
	"github.com/stretchr/testify/mock"
)

type documentTypeRepositoryMock struct {
	mock.Mock
}

func NewDocumentTypeRepositoryMock() *documentTypeRepositoryMock {
	return new(documentTypeRepositoryMock)
}

func (d *documentTypeRepositoryMock) FetchAll(c context.Context) (result []domain.DocumentType, err error) {
	ret := d.Called(c)

	if rf, ok := ret.Get(0).(func(context.Context) []domain.DocumentType); ok {
		result = rf(c)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).([]domain.DocumentType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		err = rf(c)
	} else {
		err = ret.Error(1)
	}

	return
}

func (d *documentTypeRepositoryMock) FetchByID(c context.Context, id string) (result domain.DocumentType, err error) {
	ret := d.Called(c, id)

	if rf, ok := ret.Get(0).(func(context.Context, string) domain.DocumentType); ok {
		result = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).(domain.DocumentType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		err = rf(c, id)
	} else {
		err = ret.Error(1)
	}

	return
}
