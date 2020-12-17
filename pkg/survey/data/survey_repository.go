package data

import (
	"context"
	"ephelsa/my-career/pkg/survey/domain"
	"github.com/gofiber/fiber/v2"
)

type SurveyLocalRepository interface {
	FetchAll(ctx context.Context) ([]domain.Survey, error)
	FetchActiveSurveyById(ctx context.Context, surveyId string) (domain.SurveyWithQuestions, error)
	NewQuestionAnswer(ctx context.Context, answer domain.UserAnswer) error
}

type SurveyServerRepository interface {
	// FetchAll is http.MethodGet
	FetchAll(c *fiber.Ctx) error
	// FetchActiveSurveyById is http.MethodGet
	FetchActiveSurveyById(c *fiber.Ctx) error
	// NewQuestionAnswer is http.MethodPut
	NewQuestionAnswer(c *fiber.Ctx) error
}
