package mock

import (
	"context"
	"ephelsa/my-career/pkg/location/domain"
	"github.com/stretchr/testify/mock"
)

type locationRepositoryMock struct {
	mock.Mock
}

func NewLocationRepositoryMock() *locationRepositoryMock {
	return new(locationRepositoryMock)
}

func (r *locationRepositoryMock) FetchAllCountries(c context.Context) (result []domain.Country, err error) {
	ret := r.Called(c)

	if rf, ok := ret.Get(0).(func(context.Context) []domain.Country); ok {
		result = rf(c)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).([]domain.Country)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		err = rf(c)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *locationRepositoryMock) FetchDepartmentsByCountry(c context.Context, countryCode string) (result []domain.Department, err error) {
	ret := r.Called(c, countryCode)

	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Department); ok {
		result = rf(c, countryCode)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).([]domain.Department)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		err = rf(c, countryCode)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *locationRepositoryMock) FetchMunicipalitiesByDepartmentAndCountry(c context.Context, countryCode string, departmentCode string) (result []domain.Municipality, err error) {
	ret := r.Called(c, countryCode, departmentCode)

	if rf, ok := ret.Get(0).(func(context.Context, string, string) []domain.Municipality); ok {
		result = rf(c, countryCode, departmentCode)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).([]domain.Municipality)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		err = rf(c, countryCode, departmentCode)
	} else {
		err = ret.Error(1)
	}

	return
}
