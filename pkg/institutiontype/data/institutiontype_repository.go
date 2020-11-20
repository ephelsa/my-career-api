package data

import (
	"context"
	"ephelsa/my-career/pkg/institutiontype/domain"
)

type InstitutionTypeRepository interface {
	// FetchAll returns an array of domain.InstitutionType
	FetchAll(c context.Context) ([]domain.InstitutionType, error)
}
