package server

import (
	"ephelsa/my-career/pkg/documenttype/data"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}

func (h *handler) FetchByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.Repository.FetchByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}
