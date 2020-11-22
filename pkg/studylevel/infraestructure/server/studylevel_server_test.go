package server

import (
	"encoding/json"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/shared/infrastructure/server"
	"ephelsa/my-career/pkg/studylevel/domain"
	testMock "ephelsa/my-career/test/database/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandler_FetchAllHandler(t *testing.T) {
	tests := []struct {
		description string

		httpMethod string
		route      string

		expectedCode  int
		expectedBody  sharedDomain.Response
		expectedError error
	}{
		{
			description:  "fetch all study levels",
			httpMethod:   http.MethodGet,
			route:        "/study-level/",
			expectedCode: http.StatusOK,
			expectedBody: sharedDomain.SuccessResponse([]domain.StudyLevel{
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
			description:  "fetch all study levels getting an error",
			httpMethod:   http.MethodGet,
			route:        "/study-level/",
			expectedCode: http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: sharedDomain.UnExpectedError,
				Details: sharedDomain.ResourcesEmpty.Error(),
			}),
			expectedError: sharedDomain.ResourcesEmpty,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, err := http.NewRequest(test.httpMethod, test.route, nil)
			assert.NoErrorf(t, err, test.description)

			mockRepo := testMock.NewStudyLevelRepositoryMock()
			mockRepo.On("FetchAll", mock.Anything).Return(test.expectedBody.Result, test.expectedError)

			app := server.NewServer().Server
			NewStudyLevelServer(app, mockRepo)
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
