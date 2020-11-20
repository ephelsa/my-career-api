package server

import (
	"ephelsa/my-career/pkg/documenttype/data"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type handler struct {
	Local data.DocumentTypeLocalRepository
}

func NewDocumentTypeServer(remote *fiber.App, repo data.DocumentTypeLocalRepository) {
	handler := &handler{
		Local: repo,
	}

	dt := remote.Group("/document-type")
	dt.Get("/", handler.FetchAll)
	dt.Get("/:id", handler.FetchByID)
}

func (h *handler) FetchAll(c *fiber.Ctx) error {
	result, err := h.Local.FetchAll(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: "An error occurs fetching all document types",
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}

func (h *handler) FetchByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.Local.FetchByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: fmt.Sprintf("An error occurs fetching by %s", id),
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}
