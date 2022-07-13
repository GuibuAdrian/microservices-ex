package config

import "os"

const apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"

var gihubAccessToken = os.Getenv(apiGithubAccessToken)

func GetGitHubAccessToken() string {
	return gihubAccessToken
}
