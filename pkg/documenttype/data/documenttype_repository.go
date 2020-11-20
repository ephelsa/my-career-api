package data

import (
	"context"
	"ephelsa/my-career/pkg/documenttype/domain"
)

type DocumentTypeLocalRepository interface {
	// FetchAll fetch all domain.DocumentType
	FetchAll(c context.Context) ([]domain.DocumentType, error)
	// FetchByID fetch a domain.DocumentType searching by an id
	FetchByID(c context.Context, id string) (domain.DocumentType, error)
}
