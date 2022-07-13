package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	restclient "microservices-ex-app/src/api/client/rest"
	"microservices-ex-app/src/api/domain/github"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func createAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHErrorResponse) {
	header := createAuthorizationHeader(accessToken)

	headers := http.Header{}
	headers.Set(headerAuthorization, header)

	response, err := restclient.Post(urlCreateRepo, request, headers)
	fmt.Println(response)
	fmt.Println(err)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s\n", err.Error()))
		return nil, &github.GitHErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitHErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GitHErrorResponse
		if err2 := json.Unmarshal(bytes, &errResponse); err2 != nil {
			return nil, &github.GitHErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s\n", err.Error()))
		return nil, &github.GitHErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response"}
	}

	return &result, nil
}
