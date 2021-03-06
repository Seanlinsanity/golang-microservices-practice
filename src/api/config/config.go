package config

import (
	"log"
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	log.Printf("secret: %s, access token: %s", apiGithubAccessToken, githubAccessToken)
	return githubAccessToken
}
