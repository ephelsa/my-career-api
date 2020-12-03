package server

import (
	"encoding/json"
	"ephelsa/my-career/pkg/location/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/shared/infrastructure/server"
	testMock "ephelsa/my-career/test/database/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandler_FetchAllCountries(t *testing.T) {
	tests := []struct {
		description string

		httpMethod string
		route      string

		expectedCode      int
		expectedBody      sharedDomain.Response
		expectedErrorType error
	}{
		{
			description:  "success countries",
			httpMethod:   http.MethodGet,
			route:        "/location/country",
			expectedCode: http.StatusOK,
			expectedBody: sharedDomain.SuccessResponse([]domain.Country{
				{
					ISOCode: "CO",
					Name:    "Colombia",
				},
				{
					ISOCode: "MX",
					Name:    "Mexico",
				},
			}),
			expectedErrorType: nil,
		},
		{
			description:  "empty resource error",
			httpMethod:   http.MethodGet,
			route:        "/location/country",
			expectedCode: http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: sharedDomain.UnexpectedError,
				Details: sharedDomain.ResourcesEmpty.Error(),
			}),
			expectedErrorType: sharedDomain.ResourcesEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			mockRepo := testMock.NewLocationRepositoryMock()
			mockRepo.On("FetchAllCountries", mock.Anything).Return(tt.expectedBody.Result, tt.expectedErrorType).Once()

			req, err := http.NewRequest(tt.httpMethod, tt.route, nil)
			assert.NoErrorf(t, err, tt.description)

			app := server.NewServer().Server
			NewLocationServer(app, mockRepo)
			res, err := app.Test(req, -1)
			assert.NoErrorf(t, err, tt.description)

			assert.Equalf(t, tt.expectedCode, res.StatusCode, tt.description)

			body, err := ioutil.ReadAll(res.Body)
			assert.NoErrorf(t, err, tt.description)

			expBody, err := json.Marshal(tt.expectedBody)
			assert.NoErrorf(t, err, tt.description)

			assert.Equalf(t, string(expBody), string(body), tt.description)
		})
	}
}

func TestHandler_FetchAllDepartmentsByCountry(t *testing.T) {
	tests := []struct {
		description string

		httpMethod string
		route      string

		countryArg string

		expectedCode      int
		expectedBody      sharedDomain.Response
		expectedErrorType error
	}{
		{
			description:  "success departments by country code",
			httpMethod:   http.MethodGet,
			route:        "/location/country/CO/department",
			countryArg:   "CO",
			expectedCode: http.StatusOK,
			expectedBody: sharedDomain.SuccessResponse([]domain.Department{
				{
					CountryCode:    "CO",
					DepartmentCode: "05",
					Name:           "Antioquia",
				},
				{
					CountryCode:    "CO",
					DepartmentCode: "70",
					Name:           "Sincelejo",
				},
			}),
			expectedErrorType: nil,
		},
		{
			description:  "error empty resource",
			httpMethod:   http.MethodGet,
			route:        "/location/country/CO/department",
			countryArg:   "CO",
			expectedCode: http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: sharedDomain.UnexpectedError,
				Details: sharedDomain.ResourcesEmpty.Error(),
			}),
			expectedErrorType: sharedDomain.ResourcesEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			mockRepo := testMock.NewLocationRepositoryMock()
			mockRepo.On("FetchDepartmentsByCountry", mock.Anything, tt.countryArg).Return(tt.expectedBody.Result, tt.expectedErrorType).Once()

			req, err := http.NewRequest(tt.httpMethod, tt.route, nil)
			assert.NoErrorf(t, err, tt.description)

			app := server.NewServer().Server
			NewLocationServer(app, mockRepo)
			res, err := app.Test(req, -1)
			assert.NoErrorf(t, err, tt.description)

			assert.Equalf(t, tt.expectedCode, res.StatusCode, tt.description)

			body, err := ioutil.ReadAll(res.Body)
			assert.NoErrorf(t, err, tt.description)

			expBody, err := json.Marshal(tt.expectedBody)
			assert.NoErrorf(t, err, tt.description)

			assert.Equalf(t, string(expBody), string(body), tt.description)
		})
	}
}

func TestHandler_FetchAllMunicipalitiesByDepartmentAndCountry(t *testing.T) {
	tests := []struct {
		description string

		httpMethod string
		route      string

		countryArg    string
		departmentArg string

		expectedCode      int
		expectedBody      sharedDomain.Response
		expectedErrorType error
	}{
		{
			description:   "success response",
			httpMethod:    http.MethodGet,
			route:         "/location/country/CO/department/05/municipality",
			countryArg:    "CO",
			departmentArg: "05",
			expectedCode:  http.StatusOK,
			expectedBody: sharedDomain.SuccessResponse([]domain.Municipality{
				{
					CountryCode:      "CO",
					DepartmentCode:   "05",
					MunicipalityCode: "01",
					Name:             "Medellin",
				},
				{
					CountryCode:      "CO",
					DepartmentCode:   "05",
					MunicipalityCode: "227",
					Name:             "Envigado",
				},
			}),
			expectedErrorType: nil,
		},
		{
			description:   "empty resource error",
			httpMethod:    http.MethodGet,
			route:         "/location/country/CO/department/05/municipality",
			countryArg:    "CO",
			departmentArg: "05",
			expectedCode:  http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: sharedDomain.UnexpectedError,
				Details: sharedDomain.ResourcesEmpty.Error(),
			}),
			expectedErrorType: sharedDomain.ResourcesEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			mockRepo := testMock.NewLocationRepositoryMock()
			mockRepo.On("FetchMunicipalitiesByDepartmentAndCountry", mock.Anything, tt.countryArg, tt.departmentArg).Return(tt.expectedBody.Result, tt.expectedErrorType)

			req, err := http.NewRequest(tt.httpMethod, tt.route, nil)
			assert.NoErrorf(t, err, tt.description)

			app := server.NewServer().Server
			NewLocationServer(app, mockRepo)
			res, err := app.Test(req, -1)
			assert.NoErrorf(t, err, tt.description)

			assert.Equalf(t, tt.expectedCode, res.StatusCode, tt.description)

			body, err := ioutil.ReadAll(res.Body)
			assert.NoErrorf(t, err, tt.description)

			expBody, err := json.Marshal(tt.expectedBody)
			assert.NoErrorf(t, err, tt.description)

			assert.Equalf(t, string(expBody), string(body), tt.description)
		})
	}
}
