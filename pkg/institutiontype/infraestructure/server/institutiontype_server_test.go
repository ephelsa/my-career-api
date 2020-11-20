package server

import (
	"encoding/json"
	"ephelsa/my-career/pkg/institutiontype/data"
	"ephelsa/my-career/pkg/institutiontype/domain"
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

		repository data.InstitutionTypeRepository

		expectedCode  int
		expectedBody  sharedDomain.Response
		expectedError bool
	}{
		{
			description:  "fetch all institution type",
			httpMethod:   "GET",
			route:        "/institution-type/",
			repository:   mock.FakeInstitutionTypeFullData(),
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
			expectedError: false,
		},
		{
			description:  "fetch all institution type with errors",
			httpMethod:   "GET",
			route:        "/institution-type/",
			repository:   mock.FakeInstitutionTypeErrorData(),
			expectedCode: http.StatusInternalServerError,
			expectedBody: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: "An error occurs fetch all institution types",
				Details: "something wrong retrieving resource",
			}),
			expectedError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, _ := http.NewRequest(test.httpMethod, test.route, nil)
			app := server.NewServer().Server
			NewInstitutionTypeServer(app, test.repository)
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
