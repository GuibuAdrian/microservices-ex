package config

import "os"

const (
	secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel                = "info"
	goEnvironment           = "GO_ENVIRONMENT"
	production              = "production"
)

var gihubAccessToken = os.Getenv(secretGithubAccessToken)

func GetGitHubAccessToken() string {
	return gihubAccessToken
}

func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
