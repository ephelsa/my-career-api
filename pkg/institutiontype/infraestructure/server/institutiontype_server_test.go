package server

import (
	"encoding/json"
	"ephelsa/my-career/pkg/institutiontype/domain"
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
			description:  "fetch all institution type",
			httpMethod:   http.MethodGet,
			route:        "/institution-type/",
			expectedCode: http.StatusOK,
			expectedBody: sharedDomain.SuccessResponse([]domain.InstitutionType{
				{
					Id:   0,
					Name: "Zero",
				},
				{
					Id:   1,
					Name: "One",
				},
				{
					Id:   2,
					Name: "Two",
				},
			}),
			expectedError: nil,
		},
		{
			description:  "fetch all institution type with errors",
			httpMethod:   http.MethodGet,
			route:        "/institution-type/",
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

			mockRepo := testMock.NewInstitutionTypeRepositoryMock()
			mockRepo.On("FetchAll", mock.Anything).Return(test.expectedBody.Result, test.expectedError)

			app := server.NewServer().Server
			NewInstitutionTypeServer(app, mockRepo)
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
