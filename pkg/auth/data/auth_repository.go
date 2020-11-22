package data

import (
	"context"
	"ephelsa/my-career/pkg/auth/domain"
)

type AuthRepository interface {
	// IsUserRegistered verify if the possible new user exists
	IsUserRegistered(c context.Context, email string) (res bool, err error)
	// IsUserRegistryConfirmed verify if a previous user has confirm the registry
	IsUserRegistryConfirmed(c context.Context, email string) (res bool, err error)
	// Register create a new user register
	Register(c context.Context, r domain.Register) (res domain.RegisterSuccess, err error)
	// IsAuthSuccess verify if domain.AuthCredentials match
	IsAuthSuccess(c context.Context, auth domain.AuthCredentials) (res bool, err error)
}
