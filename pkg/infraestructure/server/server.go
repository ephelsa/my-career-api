package server

import (
	"ephelsa/my-career/pkg/infraestructure/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Server struct {
	App *fiber.App
	DB  *database.Information
}

func New(data *database.Information) *Server {
	return &Server{
		App: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: true,
		}),
		DB: data,
	}
}

func (s *Server) Start(port string) {
	log.Fatal(s.App.Listen(port))
}

func (s *Server) Close() {
	if err := s.App.Shutdown(); err != nil {
		log.Panic(err)
	}
}

func (s *Server) Router() {
	type DocumentType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	s.App.Get("/document-type/all", func(c *fiber.Ctx) (err error) {
		query := `SELECT id, value FROM document_type;`
		rows, err := s.DB.Instance.Query(query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer func() {
			err = rows.Close()
		}()
		var result []DocumentType

		for rows.Next() {
			dt := DocumentType{}
			if err := rows.Scan(&dt.ID, &dt.Name); err != nil {
				return err
			}

			result = append(result, dt)
		}

		return c.JSON(result)
	})
}
