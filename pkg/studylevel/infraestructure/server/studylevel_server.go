package server

import (
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/studylevel/data"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type handler struct {
	Repository data.StudyLevelRepository
}

func NewStudyLevelServer(remote *fiber.App, repo data.StudyLevelRepository) {
	handler := &handler{
		Repository: repo,
	}

	studyLevel := remote.Group("/study-level")
	studyLevel.Get("/", handler.FetchAll)
}

func (h *handler) FetchAll(c *fiber.Ctx) error {
	result, err := h.Repository.FetchAll(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}
