package data

import (
	"context"
	"ephelsa/my-career/pkg/survey/domain"
	"github.com/gofiber/fiber/v2"
)

type SurveyLocalRepository interface {
	FetchAll(ctx context.Context, user domain.UserAnswer) ([]domain.Survey, error)
	FetchActiveSurveyById(ctx context.Context, surveyId string) (domain.SurveyWithQuestions, error)
	NewQuestionAnswer(ctx context.Context, answer domain.UserAnswer) error
	BulkQuestionAnswer(ctx context.Context, answers []*domain.UserAnswer) error
}

type SurveyServerRepository interface {
	// FetchAll is http.MethodPost
	FetchAll(c *fiber.Ctx) error
	// FetchActiveSurveyById is http.MethodGet
	FetchActiveSurveyById(c *fiber.Ctx) error
	// NewQuestionAnswer is http.MethodPut
	NewQuestionAnswer(c *fiber.Ctx) error
	// BulkQuestionAnswer is http.MethodPost
	BulkQuestionAnswer(c *fiber.Ctx) error
}
