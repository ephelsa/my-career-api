package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

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
