package server

import (
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
	"ephelsa/my-career/pkg/user/data"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	Repository data.UserLocalRepository
}

func NewUserServer(remote *fiber.App, repository data.UserLocalRepository) data.UserServerRepository {
	handler := &handler{
		Repository: repository,
	}

	user := remote.Group("/user")
	user.Get("/:email", handler.InformationByEmail)

	return handler
}

func (h *handler) InformationByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	result, err := h.Repository.InformationByEmail(c.Context(), email)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, result)
}
