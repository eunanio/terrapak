package github

import (
	"encoding/json"
	"fmt"
	"terrapak/internal/api/auth/types"
	"terrapak/internal/config"

	"net/http"

	"github.com/gin-gonic/gin"
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
// You can use this method as a common way to validate claims.
func (g GithubProvider) PostAuth(token string, c *gin.Context) {
	gc := config.GetDefault()
	if gc.AuthProvider.Organization != "" {
		user, err := g.UserInfo(token); if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/orgs/%s/members/%s", gc.AuthProvider.Organization, user.Login), nil); if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		client := &http.Client{}
		resp, err := client.Do(req); if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		if resp.StatusCode != 204 {
			fmt.Println(err)
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		
	}

}

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

	return userInfo,nil
}

func (GithubProvider) UserEmail(token string) (string,error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil); if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req); if err != nil {
		fmt.Println(err)
		return	"", err
	}
	defer resp.Body.Close()

	var userEmails []types.UserEmail
	if err := json.NewDecoder(resp.Body).Decode(&userEmails); err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, email := range userEmails {
		if email.Primary && email.Verified {
			return email.Email, nil
		}
	}

	return "", fmt.Errorf("no primary email found")
}

