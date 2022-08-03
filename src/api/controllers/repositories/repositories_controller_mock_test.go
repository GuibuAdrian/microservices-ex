package repositories

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"microservices-ex-app/src/api/domain/repositories"
	"microservices-ex-app/src/api/services"
	"microservices-ex-app/src/api/utils/errors"
	"microservices-ex-app/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	funcCreateRepo  func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
)

type repoServiceMock struct {
}

func (s *repoServiceMock) CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo("", request)
}
func (s *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	return funcCreateRepos(request)
}

//This test mocks the repoService. Isolates the controller and does not invoke repositories_service nor restclient
func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Id:    321,
			Owner: "golang",
			Name:  "mocked service",
		}, nil
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := test_utils.GetMockedContext(request, response)

	c.Request = request

	CreateRepo(c)
	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 321, result.Id)
	assert.EqualValues(t, "mocked service", result.Name)
	assert.EqualValues(t, "golang", result.Owner)
}

func TestCreateRepoErrorFromGitHubMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := test_utils.GetMockedContext(request, response)

	c.Request = request

	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	/*
		var result errors.ApiError
		err := json.Unmarshal(, &result)
		apiErr, err := errors.NewApiErrorWithBytes(response.Body.Bytes())
		assert.Nil(t, err)
		assert.NotNil(t, apiErr)
		assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
		assert.EqualValues(t, "invalid repository name", apiErr.Message())
	*/
}
