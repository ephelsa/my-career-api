package server

import (
	"ephelsa/my-career/pkg/location/data"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type handler struct {
	Repository data.LocationRepository
}

func NewLocationServer(remote *fiber.App, repo data.LocationRepository) {
	handler := &handler{
		Repository: repo,
	}

	location := remote.Group("/location")
	location.Get("/country", handler.FetchAllCountries)
	location.Get("/country/:countryCode/department", handler.FetchAllDepartmentsByCountry)
	location.Get("/country/:countryCode/department/:departmentCode/municipality", handler.FetchAllMunicipalitiesByDepartmentAndCountry)
}

// FetchAllCountries provide all countries
func (h *handler) FetchAllCountries(c *fiber.Ctx) error {
	result, err := h.Repository.FetchAllCountries(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}

// FetchAllDepartmentsByCountry provide all domain.Department by domain.Country
func (h *handler) FetchAllDepartmentsByCountry(c *fiber.Ctx) error {
	countryCode := c.Params("countryCode")
	result, err := h.Repository.FetchDepartmentsByCountry(c.Context(), countryCode)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}

// FetchAllMunicipalitiesByDepartmentAndCountry provide all domain.Municipality searching by domain.Country and
// domain.Department
func (h *handler) FetchAllMunicipalitiesByDepartmentAndCountry(c *fiber.Ctx) error {
	countryCode := c.Params("countryCode")
	departmentCode := c.Params("departmentCode")
	result, err := h.Repository.FetchMunicipalitiesByDepartmentAndCountry(c.Context(), countryCode, departmentCode)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(sharedDomain.ErrorResponse(sharedDomain.Error{
			Message: sharedDomain.UnExpectedError,
			Details: err.Error(),
		}))
	}

	return c.JSON(sharedDomain.SuccessResponse(result))
}
