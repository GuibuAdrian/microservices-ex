package config

import "os"

const secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"

var gihubAccessToken = os.Getenv(secretGithubAccessToken)

func GetGitHubAccessToken() string {
	return gihubAccessToken
}
