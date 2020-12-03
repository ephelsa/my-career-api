package server

import (
	"ephelsa/my-career/pkg/documenttype/data"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	Repository data.DocumentTypeRepository
}

func NewDocumentTypeServer(remote *fiber.App, repo data.DocumentTypeRepository) {
	handler := &handler{
		Repository: repo,
	}

	dt := remote.Group("/document-type")
	dt.Get("/", handler.FetchAll)
	dt.Get("/:id", handler.FetchByID)
}

func (h *handler) FetchAll(c *fiber.Ctx) error {
	result, err := h.Repository.FetchAll(c.Context())
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, result)
}

func (h *handler) FetchByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.Repository.FetchByID(c.Context(), id)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, result)
}
