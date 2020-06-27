package services

import (
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/config"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/domain/github"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/domain/repositories"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/providers/github_provider"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
	"strings"
)

type repoService struct {
}

type RepoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService RepoServiceInterface
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

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
