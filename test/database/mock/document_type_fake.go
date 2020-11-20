package mock

import (
	"context"
	"ephelsa/my-career/pkg/documenttype/data"
	"ephelsa/my-career/pkg/documenttype/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"fmt"
)

type fullData struct{}

type errorData struct{}

// FakeFullData returns fake data
func FakeFullData() data.DocumentTypeLocalRepository {
	return &fullData{}
}

func (l *fullData) FetchAll(_ context.Context) ([]domain.DocumentType, error) {
	return []domain.DocumentType{
		{
			Id:   "1",
			Name: "First",
		},
		{
			Id:   "2",
			Name: "Second",
		},
		{
			Id:   "3",
			Name: "Third",
		},
	}, nil
}

func (l *fullData) FetchByID(_ context.Context, id string) (domain.DocumentType, error) {
	return domain.DocumentType{
		Id:   id,
		Name: "Data",
	}, nil
}

// FakeErrorData returns error in methods
func FakeErrorData() data.DocumentTypeLocalRepository {
	return &errorData{}
}

func (l *errorData) FetchAll(_ context.Context) ([]domain.DocumentType, error) {
	return []domain.DocumentType{}, sharedDomain.ResourcesEmpty
}

func (l *errorData) FetchByID(_ context.Context, id string) (domain.DocumentType, error) {
	return domain.DocumentType{}, fmt.Errorf("something wrong retrieving data with id %s", id)
}
