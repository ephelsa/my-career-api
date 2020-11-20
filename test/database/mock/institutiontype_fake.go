package mock

import (
	"context"
	"ephelsa/my-career/pkg/institutiontype/data"
	"ephelsa/my-career/pkg/institutiontype/domain"
	"fmt"
)

type institutionTypeFullData struct{}

type institutionTypeErrorData struct{}

// FakeInstitutionTypeFullData returns fake data
func FakeInstitutionTypeFullData() data.InstitutionTypeRepository {
	return &institutionTypeFullData{}
}

func (i *institutionTypeFullData) FetchAll(_ context.Context) ([]domain.InstitutionType, error) {
	return []domain.InstitutionType{
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

// FakeInstitutionTypeErrorData returns error in methods
func FakeInstitutionTypeErrorData() data.InstitutionTypeRepository {
	return &institutionTypeErrorData{}
}

func (i *institutionTypeErrorData) FetchAll(_ context.Context) ([]domain.InstitutionType, error) {
	return []domain.InstitutionType{}, fmt.Errorf("something wrong retrieving resource")
}
