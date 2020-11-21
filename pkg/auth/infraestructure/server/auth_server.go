package server

import (
	"encoding/json"
	"ephelsa/my-career/pkg/auth/data"
	"ephelsa/my-career/pkg/auth/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	AuthRepository data.AuthRepository
}

func NewAuthServer(remote *fiber.App, authRepo data.AuthRepository) {
	handler := &handler{
		AuthRepository: authRepo,
	}

	auth := remote.Group("/auth")
	auth.Post("/registry", handler.Registry)
}

func (h *handler) Registry(c *fiber.Ctx) error {
	reg := &domain.Register{}
	err := json.Unmarshal(c.Body(), &reg)
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: "Something wrong",
			Details: err.Error(),
		}))
	}

	if len(reg.Password) < domain.PasswordMinLen {
		logrus.Error(err)
		return c.Status(http.StatusLengthRequired).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: domain.PasswordWithoutMinLen(domain.PasswordMinLen).Error(),
			Details: fmt.Sprintf("The current password has a length of %d", len(reg.Password)),
		}))
	}

	exists, err := h.AuthRepository.IsUserRegistered(c.Context(), reg.Email)
	if exists {
		err = domain.UnRegisteredUser(reg.Email)
		logrus.Error(err)
		return c.Status(http.StatusNotAcceptable).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: "User is registered",
			Details: err.Error(),
		}))
	}

	res, err := h.AuthRepository.Register(c.Context(), *reg)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: "Something wrong",
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(res))
}
