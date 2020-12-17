package data

import (
	"context"
	authDomain "ephelsa/my-career/pkg/auth/domain"
	"ephelsa/my-career/pkg/user/domain"
	"github.com/gofiber/fiber/v2"
)

type UserLocalRepository interface {
	InformationByEmail(ctx context.Context, email string) (domain.User, error)
	StoreUserInformation(ctx context.Context, user domain.User) (authDomain.RegisterSuccess, error)
}

type UserServerRepository interface {
	InformationByEmail(c *fiber.Ctx) error
}
