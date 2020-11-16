package documenttype

import "github.com/gofiber/fiber/v2"

type RemoteRepository struct {
	Repository Repository
}

func (rr *RemoteRepository) FetchAllHandler(c *fiber.Ctx) error {
	result, err := rr.Repository.FetchAll(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(result)
}
