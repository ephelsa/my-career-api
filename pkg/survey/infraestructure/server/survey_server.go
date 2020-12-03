package server

import (
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
	"ephelsa/my-career/pkg/survey/data"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	Repository data.SurveyRepository
}

func NewSurveyServer(remote *fiber.App, repo data.SurveyRepository) data.SurveyServerRepository {
	handler := &handler{
		Repository: repo,
	}

	survey := remote.Group("/survey")
	survey.Get("/", handler.FetchAll)
	survey.Get("/:id/questions-with-answers", handler.FetchActiveSurveyById)

	return handler
}

func (h *handler) FetchAll(c *fiber.Ctx) error {
	result, err := h.Repository.FetchAll(c.Context())
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	if len(result) == 0 {
		return sharedServer.NotFound(c, sharedDomain.Error{
			Message: sharedDomain.ResourceEmpty,
			Details: sharedDomain.ResourcesEmpty.Error(),
		})
	}

	return sharedServer.OK(c, result)
}

func (h *handler) FetchActiveSurveyById(c *fiber.Ctx) error {
	surveyId := c.Params("id")
	result, err := h.Repository.FetchActiveSurveyById(c.Context(), surveyId)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	if len(result.Questions) == 0 {
		return sharedServer.NotFound(c, sharedDomain.Error{
			Message: sharedDomain.ResourceEmpty,
			Details: sharedDomain.ResourcesEmpty.Error(),
		})
	}

	return sharedServer.OK(c, result)
}
