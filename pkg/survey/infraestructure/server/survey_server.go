package server

import (
	"encoding/json"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
	"ephelsa/my-career/pkg/survey/data"
	"ephelsa/my-career/pkg/survey/domain"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type handler struct {
	Repository data.SurveyLocalRepository
}

func NewSurveyServer(remote *fiber.App, repo data.SurveyLocalRepository) data.SurveyServerRepository {
	handler := &handler{
		Repository: repo,
	}

	survey := remote.Group("/survey")
	survey.Post("/by-user", handler.FetchAll)
	survey.Get("/:id/questions-with-answers", handler.FetchActiveSurveyById)
	survey.Put("/:id/answer", handler.NewQuestionAnswer)

	return handler
}

func (h *handler) FetchAll(c *fiber.Ctx) error {
	body := c.Body()
	userInfo := domain.UserAnswer{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	result, err := h.Repository.FetchAll(c.Context(), userInfo)
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

func (h *handler) NewQuestionAnswer(c *fiber.Ctx) error {
	surveyId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}
	body := c.Body()
	userAnswer := domain.UserAnswer{}
	if err := json.Unmarshal(body, &userAnswer); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}
	userAnswer.Survey = surveyId

	if err := h.Repository.NewQuestionAnswer(c.Context(), userAnswer); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, nil)
}
