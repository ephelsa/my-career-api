package data

import (
	"context"
	"ephelsa/my-career/pkg/survey/domain"
	"github.com/gofiber/fiber/v2"
)

type SurveyRepository interface {
	FetchAll(ctx context.Context) ([]domain.Survey, error)
	FetchActiveSurveyById(ctx context.Context, surveyId string) (domain.SurveyWithQuestions, error)
}

type SurveyServerRepository interface {
	FetchAll(c *fiber.Ctx) error
	FetchActiveSurveyById(c *fiber.Ctx) error
}
