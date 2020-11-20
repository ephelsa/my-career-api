package server

import (
	"ephelsa/my-career/pkg/institutiontype/data"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: "An error occurs fetch all institution types",
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}
