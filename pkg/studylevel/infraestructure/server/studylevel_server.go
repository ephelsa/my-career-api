package server

import (
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
	"ephelsa/my-career/pkg/studylevel/data"
	"github.com/gofiber/fiber/v2"
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
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, result)
}
