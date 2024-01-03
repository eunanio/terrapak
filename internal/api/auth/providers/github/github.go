package github

import (
	"encoding/json"
	"fmt"
	"terrapak/internal/api/auth/types"
	"terrapak/internal/config"

	"net/http"

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

func (GithubProvider) Authenticate(token string) {}

func (GithubProvider) UserInfo(token string) (types.UserInfo,error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil); if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req); if err != nil {
		fmt.Println(err)
		return	types.UserInfo{}, err
	}
	defer resp.Body.Close()

	var userInfo types.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		fmt.Println(err)
		return userInfo, err
	}

	fmt.Println(userInfo)

	return userInfo,nil
}