package server

import (
	"encoding/json"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/shared/infrastructure/server"
	"ephelsa/my-career/pkg/studylevel/data"
	"ephelsa/my-career/pkg/studylevel/domain"
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

		repository data.StudyLevelRepository

		expectedCode  int
		expectedBody  sharedDomain.Response
		expectedError bool
	}{
		{
			description:  "fetch all study levels",
			httpMethod:   "GET",
			route:        "/study-level/",
			repository:   mock.FakeStudyLevelFullData(),
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
			expectedError: false,
		},
		{
			description:  "fetch all study levels getting an error",
			httpMethod:   "GET",
			route:        "/study-level/",
			repository:   mock.FakeStudyLevelErrorData(),
			expectedCode: http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: "An error occurs fetching all study levels",
				Details: "resource is empty",
			}),
			expectedError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, _ := http.NewRequest(test.httpMethod, test.route, nil)
			app := server.NewServer().Server
			NewStudyLevelServer(app, test.repository)
			res, err := app.Test(req, -1)

			assert.Equalf(t, test.expectedError, err != nil, test.description)

			//goland:noinspection GoNilness
			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

			//goland:noinspection GoNilness
			body, err := ioutil.ReadAll(res.Body)
			expBody, _ := json.Marshal(test.expectedBody)

			assert.Nilf(t, err, test.description)
			assert.Equalf(t, string(expBody), string(body), test.description)
		})
	}
}
