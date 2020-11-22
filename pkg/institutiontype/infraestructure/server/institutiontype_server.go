package server

import (
	"ephelsa/my-career/pkg/institutiontype/data"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	Repository data.InstitutionTypeRepository
}

func NewInstitutionTypeServer(remote *fiber.App, repo data.InstitutionTypeRepository) {
	handler := &handler{
		Repository: repo,
	}

	institutionType := remote.Group("/institution-type")
	institutionType.Get("/", handler.FetchAll)
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
