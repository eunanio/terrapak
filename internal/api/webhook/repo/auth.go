package repo

import (
	"fmt"
	"terrapak/internal/api/webhook/rest"
	"terrapak/internal/config"

	"time"

	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
)

type InstallResponse struct {
	AccessTokenURL string `json:"access_tokens_url"`
}

type AccessCodeResponse struct {
	Token string `json:"token"`
}

func createJWT() (string,error) {
	gc := config.GetDefault()
	claims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
		"iss": gc.GitHubAppConfig.AppId,
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodRS256,claims)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(gc.GitHubAppConfig.PrivateKey)); if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	tokenString, err := token.SignedString(privateKey); if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	return tokenString, nil
}

func GetAccessToken(installationId int) (string,error) {

	endpoint := fmt.Sprintf("https://api.github.com/app/installations/%d/access_tokens", installationId)
	jwt, err := createJWT(); if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}
	client := rest.New(jwt)

	resp, err := client.Post(endpoint,"application/json",nil); if err != nil {
		return "", fmt.Errorf("error ac: %s", err)
	}
	defer resp.Body.Close()
	ac := AccessCodeResponse{}
	err = json.NewDecoder(resp.Body).Decode(&ac); if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}
	return ac.Token, nil
}