package github

import (
	"terrapak/internal/config"

	"golang.org/x/oauth2"
)

const (
	authURL  = "https://github.com/login/oauth/authorize"
	tokenURL = "https://github.com/login/oauth/access_token"
)

type GithubProvider struct {}

func New() GithubProvider {
	return GithubProvider{}
}

func (GithubProvider) Name() string {
	return "github"
}

func (g GithubProvider) Config() (conf oauth2.Config) {
	gc := config.GetDefault()
	scopes := []string{"user"}

	if gc.AuthProvider.Organization != "" {
		scopes = append(scopes, "read:org")
	}

	conf = oauth2.Config{
		ClientID:     gc.AuthProvider.ClientId,
		ClientSecret: gc.AuthProvider.ClientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}
	return conf
}

func (GithubProvider) UserInfo(token string) (string, error) {
	return "", nil
}