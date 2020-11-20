package mock

import (
	"context"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/studylevel/data"
	"ephelsa/my-career/pkg/studylevel/domain"
)

type studyLevelFullData struct{}
type studyLevelErrorData struct{}

// FakeStudyLevelFullData returns fake data
func FakeStudyLevelFullData() data.StudyLevelRepository {
	return &studyLevelFullData{}
}

func (*studyLevelFullData) FetchAll(_ context.Context) ([]domain.StudyLevel, error) {
	return []domain.StudyLevel{
		{
			Id:   0,
			Name: "Zero",
		},
		{
			Id:   1,
			Name: "One",
		},
		{
			Id:   2,
			Name: "Two",
		},
	}, nil
}

// FakeStudyLevelErrorData returns errors
func FakeStudyLevelErrorData() data.StudyLevelRepository {
	return &studyLevelErrorData{}
}

func (*studyLevelErrorData) FetchAll(_ context.Context) ([]domain.StudyLevel, error) {
	return []domain.StudyLevel{}, sharedDomain.ResourcesEmpty
}
