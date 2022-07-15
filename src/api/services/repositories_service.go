package services

import (
	"microservices-ex-app/src/api/config"
	"microservices-ex-app/src/api/domain/github"
	"microservices-ex-app/src/api/domain/repositories"
	"microservices-ex-app/src/api/log/option_b"
	githubprovider "microservices-ex-app/src/api/providers"
	"microservices-ex-app/src/api/utils/errors"
	"net/http"
	"sync"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	//option_a.Info("about to send request to external api", fmt.Sprintf("client_id:%s", clientId), "status:pending")
	option_b.Info("about to send request to external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "pending"),
		option_b.Field("authenticated", clientId != ""))
	response, err := githubprovider.CreateRepo(config.GetGitHubAccessToken(), request)
	if err != nil {
		//option_a.Error("response obtained from external api", err, fmt.Sprintf("client_id:%s", clientId), "status:error")
		option_b.Error("response obtained from external api", err,
			option_b.Field("client_id", clientId),
			option_b.Field("status", "error"),
			option_b.Field("authenticated", clientId != ""))
		apiErr := errors.NewApiError(err.StatusCode, err.Message)
		return nil, apiErr
	}

	//option_a.Info("response obtained from external api", fmt.Sprintf("client_id:%s", clientId), "status:success")
	option_b.Info("response obtained from external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "success"),
		option_b.Field("authenticated", clientId != ""))
	result := repositories.CreateRepoResponse{
		Id:    response.ID,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}

	return &result, nil
}

func (s *repoService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup //This is a control mechanism to block the execution until some work is done
	go s.handleRepoResults(&wg, input, output)

	for _, currentRequest := range requests {
		wg.Add(1) //If we send 3 request then 3 diff go routines are launched wait group has a counter of 3
		go s.createRepoConcurrent(currentRequest, input)
	}

	wg.Wait()    // This will block without continuing until de wait group reaches 0
	close(input) // Once the process releases execution we close the input to release the handleRepoResults() range iteration

	//This blocks until something arrive to the output channel and the only go routine that populates the output is handleRepoResults()
	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

//This go routine processes all the request to GitHub and then sends it to output channel
//It listens to the results generated by createRepoConcurrent() and merges them into one
func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	//Blocks the iteration until something closes input channel
	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done() // Every time we have a result coming from GitHub (success or error) we will decrease 1 unit from wg
	}
	//When the input channel closes the range of input finishes we send the results to the output
	output <- results
}

//this does not return anything. It sends the result to the given channel
func (s *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo("", input) //Sends post request to GitHub to create a repository
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}
}
