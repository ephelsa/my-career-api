package server

import (
	"context"
	"encoding/json"
	"ephelsa/my-career/pkg/auth/data"
	"ephelsa/my-career/pkg/auth/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
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
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	exists, err := h.isUserRegistered(c.Context(), reg.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	if exists {
		confirmed, err := h.isUserRegistryConfirmed(c.Context(), reg.Email)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: sharedDomain.UnExpectedError,
				Details: err.Error(),
			}))
		}

		if confirmed {
			err = data.UserRegistered(reg.Email)
			logrus.Error(err)
			return c.Status(http.StatusNotAcceptable).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: data.UserIsRegistered,
				Details: err.Error(),
			}))
		} else {
			err = data.UserRegisteredWithoutConfirm(reg.Email)
			logrus.Error(err)
			return c.Status(http.StatusNotAcceptable).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: data.UserIsRegistered,
				Details: err.Error(),
			}))
		}
	}

	if !isPasswordLenValid(reg.Password) {
		logrus.Error(err)
		return c.Status(http.StatusLengthRequired).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: data.PasswordLength,
			Details: data.PasswordWithoutMinLen(domain.PasswordMinLen).Error(),
		}))
	}

	res, err := h.AuthRepository.Register(c.Context(), *reg)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(res))
}

func (h *handler) isUserRegistered(c context.Context, email string) (bool, error) {
	exists, err := h.AuthRepository.IsUserRegistered(c, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (h *handler) isUserRegistryConfirmed(c context.Context, email string) (bool, error) {
	confirmed, err := h.AuthRepository.IsUserRegistryConfirmed(c, email)
	if err != nil {
		return false, err
	}

	return confirmed, nil
}

func isPasswordLenValid(password string) bool {
	return len(password) > domain.PasswordMinLen
}
