package server

import (
	"ephelsa/my-career/pkg/documenttype"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Server struct {
	App *fiber.App
}

func New(port string) *Server {
	server := &Server{
		App: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: true,
		}),
	}

	log.Fatal(server.App.Listen(port))

	return server
}

func (s *Server) Close() {
	if err := s.App.Shutdown(); err != nil {
		log.Panic(err)
	}
}

func (s *Server) Routes(documentTypeRepo documenttype.RemoteRepository) {
	documentType := s.App.Group("/document-type")
	documentType.Group("/all", documentTypeRepo.FetchAllHandler)
}
