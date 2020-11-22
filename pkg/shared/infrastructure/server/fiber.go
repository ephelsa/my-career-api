package server

import (
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"net/http"
)

var (
	// OK status code http.StatusOK
	OK = func(c *fiber.Ctx, r interface{}) error {
		return c.Status(http.StatusOK).JSON(sharedDomain.SuccessResponse(r))
	}
	// InternalServerError status code http.StatusInternalServerError
	InternalServerError = func(c *fiber.Ctx, err sharedDomain.Error) error {
		return errorResponse(c, http.StatusInternalServerError, err)
	}
	// NotAcceptable status code http.StatusNotAcceptable
	NotAcceptable = func(c *fiber.Ctx, err sharedDomain.Error) error {
		return errorResponse(c, http.StatusNotAcceptable, err)
	}
	// LengthRequired status code http.StatusLengthRequired
	LengthRequired = func(c *fiber.Ctx, err sharedDomain.Error) error {
		return errorResponse(c, http.StatusLengthRequired, err)
	}
	Unauthorized = func(c *fiber.Ctx, err sharedDomain.Error) error {
		return errorResponse(c, http.StatusUnauthorized, err)
	}
)

func errorResponse(c *fiber.Ctx, errorCode int, err sharedDomain.Error) error {
	return c.Status(errorCode).JSON(sharedDomain.ErrorResponse(err))
}

type Server struct {
	Server *fiber.App
}

func NewServer() *Server {
	return &Server{
		Server: fiber.New(fiber.Config{
			CaseSensitive: true,
		}),
	}
}

func (s *Server) Start(port string) {
	log.Fatal(s.Server.Listen(fmt.Sprintf(":%s", port)))
}

func (s *Server) Close() {
	if err := s.Server.Shutdown(); err != nil {
		log.Panic(err)
	}
}

func (s *Server) Middleware() {
	s.Server.Use(logger.New())
}
