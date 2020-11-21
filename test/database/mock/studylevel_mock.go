package mock

import (
	"context"
	"ephelsa/my-career/pkg/studylevel/domain"
	"github.com/stretchr/testify/mock"
)

type studyLevelRepositoryMock struct {
	mock.Mock
}

func NewStudyLevelRepositoryMock() *studyLevelRepositoryMock {
	return new(studyLevelRepositoryMock)
}

func (s *studyLevelRepositoryMock) FetchAll(c context.Context) (result []domain.StudyLevel, err error) {
	ret := s.Called(c)

	if rf, ok := ret.Get(0).(func(context.Context) []domain.StudyLevel); ok {
		result = rf(c)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).([]domain.StudyLevel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		err = rf(c)
	} else {
		err = ret.Error(1)
	}

	return
}
