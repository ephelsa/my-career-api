package mock

import (
	"context"
	"ephelsa/my-career/pkg/documenttype/data"
	"ephelsa/my-career/pkg/documenttype/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"fmt"
)

type documentTypeFullData struct{}

type documentTypeErrorData struct{}

// FakeDocumentTypeFullData returns fake data
func FakeDocumentTypeFullData() data.DocumentTypeRepository {
	return &documentTypeFullData{}
}

func (l *documentTypeFullData) FetchAll(_ context.Context) ([]domain.DocumentType, error) {
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

func (l *documentTypeFullData) FetchByID(_ context.Context, id string) (domain.DocumentType, error) {
	return domain.DocumentType{
		Id:   id,
		Name: "Data",
	}, nil
}

// FakeDocumentTypeErrorData returns error in methods
func FakeDocumentTypeErrorData() data.DocumentTypeRepository {
	return &documentTypeErrorData{}
}

func (l *documentTypeErrorData) FetchAll(_ context.Context) ([]domain.DocumentType, error) {
	return []domain.DocumentType{}, sharedDomain.ResourcesEmpty
}

func (l *documentTypeErrorData) FetchByID(_ context.Context, id string) (domain.DocumentType, error) {
	return domain.DocumentType{}, fmt.Errorf("something wrong retrieving data with id %s", id)
}
