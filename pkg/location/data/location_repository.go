package data

import (
	"context"
	"ephelsa/my-career/pkg/location/domain"
)

type LocationRepository interface {
	// FetchAllCountries fetch all domain.Country
	FetchAllCountries(c context.Context) (result []domain.Country, err error)
	// FetchDepartmentsByCountry fetch all domain.Department by domain.Country ISOCode
	FetchDepartmentsByCountry(c context.Context, countryCode string) (result []domain.Department, err error)
	// FetchMunicipalitiesByDepartment fetch all domain.Municipality by domain.Department DepartmentCode and domain.Country CountryCode
	FetchMunicipalitiesByDepartmentAndCountry(c context.Context, countryCode string, departmentCode string) (result []domain.Municipality, err error)
}
