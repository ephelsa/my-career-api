package server

import (
	"encoding/json"
	"ephelsa/my-career/pkg/documenttype/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/shared/infrastructure/server"
	testMock "ephelsa/my-career/test/database/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandler_FetchAll(t *testing.T) {
	tests := []struct {
		description string

		httpMethod string
		route      string

		expectedCode  int
		expectedBody  sharedDomain.Response
		expectedError error
	}{
		{
			description:  "fetch all document types",
			route:        "/document-type/",
			httpMethod:   http.MethodGet,
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
			expectedError: nil,
		},
		{
			description:  "fetch all document types with error",
			route:        "/document-type/",
			httpMethod:   "GET",
			expectedCode: http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: sharedDomain.UnexpectedError,
				Details: sharedDomain.ResourcesEmpty.Error(),
			}),
			expectedError: sharedDomain.ResourcesEmpty,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, err := http.NewRequest(test.httpMethod, test.route, nil)
			assert.NoErrorf(t, err, test.description)

			mockRepo := testMock.NewDocumentTypeRepositoryMock()
			mockRepo.On("FetchAll", mock.Anything).Return(test.expectedBody.Result, test.expectedError)

			app := server.NewServer().Server
			NewDocumentTypeServer(app, mockRepo)
			res, err := app.Test(req, -1)
			assert.NoErrorf(t, err, test.description)
			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

			body, err := ioutil.ReadAll(res.Body)
			assert.NoErrorf(t, err, test.description)

			expBody, err := json.Marshal(test.expectedBody)
			assert.NoErrorf(t, err, test.description)

			assert.Equalf(t, string(expBody), string(body), test.description)
		})
	}
}

func TestHandler_FetchByID(t *testing.T) {
	tests := []struct {
		description string

		httpMethod string
		route      string

		idArg string

		expectedCode  int
		expectedBody  sharedDomain.Response
		expectedError error
	}{
		{
			description:  "fetch data with id something",
			httpMethod:   "GET",
			route:        "/document-type/something",
			idArg:        "something",
			expectedCode: http.StatusOK,
			expectedBody: sharedDomain.SuccessResponse(domain.DocumentType{
				Id:   "something",
				Name: "Data",
			}),
			expectedError: nil,
		},
		{
			description:  "fetch data with id something an get an error",
			httpMethod:   "GET",
			route:        "/document-type/something",
			idArg:        "something",
			expectedCode: http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: sharedDomain.UnexpectedError,
				Details: sharedDomain.ResourcesEmpty.Error(),
			}),
			expectedError: sharedDomain.ResourcesEmpty,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, err := http.NewRequest(test.httpMethod, test.route, nil)
			assert.NoErrorf(t, err, test.description)

			mockRepo := testMock.NewDocumentTypeRepositoryMock()
			mockRepo.On("FetchByID", mock.Anything, test.idArg).Return(test.expectedBody.Result, test.expectedError)

			app := server.NewServer().Server
			NewDocumentTypeServer(app, mockRepo)
			res, err := app.Test(req, -1)
			assert.NoErrorf(t, err, test.description)
			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

			body, err := ioutil.ReadAll(res.Body)
			assert.NoErrorf(t, err, test.description)

			expBody, err := json.Marshal(test.expectedBody)
			assert.NoErrorf(t, err, test.description)
			assert.Equalf(t, string(expBody), string(body), test.description)
		})
	}
}
