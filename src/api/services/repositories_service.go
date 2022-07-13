package services

import (
	"microservices-ex-app/src/api/config"
	"microservices-ex-app/src/api/domain/github"
	"microservices-ex-app/src/api/domain/repositories"
	github_provider "microservices-ex-app/src/api/providers"
	"microservices-ex-app/src/api/utils/errors"
	"strings"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGitHubAccessToken(), request)
	if err != nil {
		apiErr := errors.NewApiError(err.StatusCode, err.Message)
		return nil, apiErr
	}

	result := repositories.CreateRepoResponse{
		Id:    response.ID,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}

	return &result, nil
}
