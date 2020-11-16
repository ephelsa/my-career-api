package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Server struct {
	App  *fiber.App
	Port string
}

func NewServer(port string) *Server {
	return &Server{
		App: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: true,
		}),
		Port: port,
	}
}

func (s *Server) Start() {
	log.Fatal(s.App.Listen(s.Port))
}
