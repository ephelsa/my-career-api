package server

import (
	"context"
	"encoding/json"
	"ephelsa/my-career/pkg/auth/data"
	"ephelsa/my-career/pkg/auth/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
	auth.Post("/login", handler.Login)
}

func (h *handler) Registry(c *fiber.Ctx) error {
	reg := &domain.Register{}
	err := json.Unmarshal(c.Body(), &reg)
	if err != nil {
		logrus.Error(err)
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		})
	}

	exists, err := h.isUserRegistered(c.Context(), reg.Email)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		})
	}

	// Verify user existence
	if exists {
		confirmed, err := h.isUserRegistryConfirmed(c.Context(), reg.Email)
		if err != nil {
			return sharedServer.InternalServerError(c, sharedDomain.Error{
				Message: sharedDomain.UnExpectedError,
				Details: err.Error(),
			})
		}

		// Verify if the user has been confirmed the email
		if confirmed {
			err = data.UserRegisteredError(reg.Email)
			logrus.Error(err)
			return sharedServer.NotAcceptable(c, sharedDomain.Error{
				Message: data.UserIsRegistered,
				Details: err.Error(),
			})
		} else {
			err = data.UserRegisteredWithoutConfirmError(reg.Email)
			logrus.Error(err)
			return sharedServer.NotAcceptable(c, sharedDomain.Error{
				Message: data.UserIsRegistered,
				Details: err.Error(),
			})
		}
	}

	// Verify password
	if !isPasswordValid(reg.Password) {
		logrus.Error(err)
		return sharedServer.LengthRequired(c, sharedDomain.Error{
			Message: data.PasswordLength,
			Details: data.PasswordWithoutMinLenError(domain.PasswordMinLen).Error(),
		})
	}

	res, err := h.AuthRepository.Register(c.Context(), *reg)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, res)
}

func (h *handler) Login(c *fiber.Ctx) error {
	cred := domain.AuthCredentials{}
	err := json.Unmarshal(c.Body(), &cred)
	if err != nil {
		logrus.Error(err)
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		})
	}

	var confirmed bool
	exists, err := h.isUserRegistered(c.Context(), cred.Email)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		})
	}

	// Verify the user existence
	if exists {
		confirmed, err = h.isUserRegistryConfirmed(c.Context(), cred.Email)
		if err != nil {
			return sharedServer.InternalServerError(c, sharedDomain.Error{
				Message: sharedDomain.UnExpectedError,
				Details: err.Error(),
			})
		}
	} else {
		return sharedServer.Unauthorized(c, sharedDomain.Error{
			Message: data.UserNotRegistered,
			Details: data.UserNotRegisteredError(cred.Email).Error(),
		})
	}

	// Verify if the user has been confirmed the account
	if confirmed {
		isIn, err := h.isAuthSuccess(c.Context(), cred)
		if err != nil {
			return sharedServer.InternalServerError(c, sharedDomain.Error{
				Message: sharedDomain.UnExpectedError,
				Details: err.Error(),
			})
		}

		// Check the credentials provided are valid
		if isIn {
			// TODO: Session token
			return sharedServer.OK(c, domain.AuthCredentials{
				Email: cred.Email,
				Token: "an incredible token will be placed here",
			})
		} else {
			return sharedServer.Unauthorized(c, sharedDomain.Error{
				Message: data.InvalidCredentials,
				Details: data.InvalidAuthError(cred.Email).Error(),
			})
		}
	} else {
		return sharedServer.NotAcceptable(c, sharedDomain.Error{
			Message: data.ConfirmEmail,
			Details: data.UserRegisteredWithoutConfirmError(cred.Email).Error(),
		})
	}
}

// isAuthSuccess verify if the credentials match with the user
func (h *handler) isAuthSuccess(c context.Context, credentials domain.AuthCredentials) (res bool, err error) {
	res, err = h.AuthRepository.IsAuthSuccess(c, credentials)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}

// isUserRegistered verify if an email has been registered previously
func (h *handler) isUserRegistered(c context.Context, email string) (bool, error) {
	exists, err := h.AuthRepository.IsUserRegistered(c, email)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return exists, nil
}

// isUserRegistryConfirmed verify if an user has been confirmed the account
func (h *handler) isUserRegistryConfirmed(c context.Context, email string) (bool, error) {
	confirmed, err := h.AuthRepository.IsUserRegistryConfirmed(c, email)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return confirmed, nil
}

// isPasswordValid verify if the password is valid
// TODO: Implement secure requirements
func isPasswordValid(password string) bool {
	return len(password) > domain.PasswordMinLen
}
