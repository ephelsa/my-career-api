package data

import (
	"context"
	"ephelsa/my-career/pkg/studylevel/domain"
)

type StudyLevelRepository interface {
	// FetchAll study levels
	FetchAll(c context.Context) ([]domain.StudyLevel, error)
}
