package server

import (
	"encoding/json"
	"ephelsa/my-career/pkg/documenttype/data"
	"ephelsa/my-career/pkg/documenttype/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/shared/infrastructure/server"
	"ephelsa/my-career/test/database/mock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandler_FetchAll(t *testing.T) {
	tests := []struct {
		description string

		httpMethod string
		route      string

		repository data.DocumentTypeLocalRepository

		expectedCode  int
		expectedBody  sharedDomain.Response
		expectedError bool
	}{
		{
			description:  "fetch all document types",
			route:        "/document-type/",
			httpMethod:   "GET",
			expectedCode: http.StatusOK,
			expectedBody: sharedDomain.SuccessResponse([]domain.DocumentType{
				{
					Id:   "1",
					Name: "First",
				},
				{
					Id:   "2",
					Name: "Second",
				},
				{
					Id:   "3",
					Name: "Third",
				},
			}),
			expectedError: false,
			repository:    mock.FakeFullData(),
		},
		{
			description:  "fetch all document types with error",
			route:        "/document-type/",
			httpMethod:   "GET",
			expectedCode: http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: "An error occurs fetching all document types",
				Details: sharedDomain.ResourcesEmpty.Error(),
			}),
			expectedError: false,
			repository:    mock.FakeErrorData(),
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(test.httpMethod, test.route, nil)
		app := server.NewServer().Server
		NewDocumentTypeServer(app, test.repository)
		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}

		//goland:noinspection GoNilness
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		//goland:noinspection GoNilness
		body, err := ioutil.ReadAll(res.Body)
		expBody, _ := json.Marshal(test.expectedBody)

		assert.Nilf(t, err, test.description)
		assert.Equalf(t, string(expBody), string(body), test.description)
	}
}

//func TestHandler_FetchByID(t *testing.T) {
//	tests := []struct {
//		description string
//
//		httpMethod string
//		route      string
//
//		repository data.DocumentTypeLocalRepository
//
//		expectedCode  int
//		expectedBody  sharedDomain.Response
//		expectedError bool
//	}{
//		{
//			description:  "fetch data with id something",
//			httpMethod:   "GET",
//			route:        "/something",
//			repository:   mock.FakeFullData(),
//			expectedCode: http.StatusOK,
//			expectedBody: sharedDomain.SuccessResponse(domain.DocumentType{
//				Id:   "something",
//				Name: "Data",
//			}),
//			expectedError: false,
//		},
//		{
//			description:  "fetch data with id something an get an error",
//			httpMethod:   "GET",
//			route:        "/something",
//			repository:   mock.FakeErrorData(),
//			expectedCode: http.StatusInternalServerError,
//			expectedBody: sharedDomain.SuccessResponse(domain.DocumentType{
//				Id:   "something",
//				Name: "Data",
//			}),
//			expectedError: false,
//		},
//	}
//
//	for _, test := range tests {
//		req, _ := http.NewRequest(test.httpMethod, test.route, nil)
//		app := fiber.New()
//	}
//}
