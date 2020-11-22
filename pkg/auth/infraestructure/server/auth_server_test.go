package server

import (
	"bytes"
	"encoding/json"
	"ephelsa/my-career/pkg/auth/data"
	"ephelsa/my-career/pkg/auth/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/shared/infrastructure/server"
	testMock "ephelsa/my-career/test/database/mock"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

type expectedAuthRepo struct {
	Result interface{}
	Error  error
}

func openFixture(filename string) (io.Reader, error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("./fixture/%s", filename))
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(file), nil
}

func TestHandler_Registry(t *testing.T) {
	tests := []struct {
		description string

		route           string
		bodyFixtureName string

		emailArg string
		passArg  string

		expectedStatus   int
		expectedResponse sharedDomain.Response

		expUserRegisteredRepo        expectedAuthRepo
		expUserRegistryConfirmedRepo expectedAuthRepo
		expRegister                  expectedAuthRepo
	}{
		{
			description:     "user exists without confirm",
			route:           "/auth/registry",
			bodyFixtureName: "Registry.json",
			emailArg:        "xephelsax@gmail.com",
			passArg:         "SuperSecretPassword",
			expectedStatus:  http.StatusNotAcceptable,
			expectedResponse: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: data.UserIsRegistered,
				Details: data.UserRegisteredWithoutConfirm("xephelsax@gmail.com").Error(),
			}),
			expUserRegisteredRepo: expectedAuthRepo{
				Result: true,
			},
			expUserRegistryConfirmedRepo: expectedAuthRepo{
				Result: false,
			},
		},
		{
			description:     "user exists with confirmed registry",
			route:           "/auth/registry",
			bodyFixtureName: "Registry.json",
			emailArg:        "xephelsax@gmail.com",
			passArg:         "SuperSecretPassword",
			expectedStatus:  http.StatusNotAcceptable,
			expectedResponse: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: data.UserIsRegistered,
				Details: data.UserRegistered("xephelsax@gmail.com").Error(),
			}),
			expUserRegisteredRepo: expectedAuthRepo{
				Result: true,
			},
			expUserRegistryConfirmedRepo: expectedAuthRepo{
				Result: true,
			},
		},
		{
			description:     "password less than min length",
			route:           "/auth/registry",
			bodyFixtureName: "RegistryWithoutMinPassword.json",
			emailArg:        "xephelsax@gmail.com",
			passArg:         "1234567",
			expectedStatus:  http.StatusLengthRequired,
			expectedResponse: sharedDomain.ErrorResponse(sharedDomain.Error{
				Message: data.PasswordLength,
				Details: data.PasswordWithoutMinLen(domain.PasswordMinLen).Error(),
			}),
			expUserRegisteredRepo: expectedAuthRepo{
				Result: false,
			},
			expUserRegistryConfirmedRepo: expectedAuthRepo{
				Result: false,
			},
			expRegister: expectedAuthRepo{
				Result: domain.RegisterSuccess{
					Email: "xephelsax@gmail.com",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			fixture, err := openFixture(tt.bodyFixtureName)
			assert.NoErrorf(t, err, tt.description)
			req, err := http.NewRequest(http.MethodPost, tt.route, fixture)
			assert.NoErrorf(t, err, tt.description)

			reg := domain.Register{}
			fixtureC, err := openFixture(tt.bodyFixtureName) // Copy
			assert.NoErrorf(t, err, tt.description)
			b, err := ioutil.ReadAll(fixtureC)
			assert.NoErrorf(t, err, tt.description)

			err = json.Unmarshal(b, &reg)
			assert.NoErrorf(t, err, tt.description)

			mockRepo := testMock.NewAuthRepositoryMock()
			mockRepo.On("IsUserRegistered", mock.Anything, tt.emailArg).Return(tt.expUserRegisteredRepo.Result, tt.expUserRegisteredRepo.Error)
			mockRepo.On("IsUserRegistryConfirmed", mock.Anything, tt.emailArg).Return(tt.expUserRegistryConfirmedRepo.Result, tt.expUserRegistryConfirmedRepo.Error)
			mockRepo.On("Register", mock.Anything, reg).Return(tt.expRegister.Result, tt.expRegister.Error)

			app := server.NewServer().Server
			NewAuthServer(app, mockRepo)
			res, err := app.Test(req, -1)
			assert.NoErrorf(t, err, tt.description)
			assert.Equalf(t, tt.expectedStatus, res.StatusCode, tt.description)

			body, err := ioutil.ReadAll(res.Body)
			assert.NoErrorf(t, err, tt.description)

			expBody, err := json.Marshal(tt.expectedResponse)
			assert.NoErrorf(t, err, tt.description)

			assert.Equalf(t, string(expBody), string(body), tt.description)
		})
	}
}
