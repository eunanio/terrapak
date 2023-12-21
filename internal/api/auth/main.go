package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

type AuthProvider interface{}

type OAuthToken struct {
	AccessToken string `json:"access_token"`
}

func Authorize(c *gin.Context) {
	//..
}

func Token(c *gin.Context) {
	//..
}

func Callback(c *gin.Context) {
	//..
}

func generateCodeVerifier() string {
    b := make([]byte, 32)
    rand.Read(b)
    return base64.RawURLEncoding.EncodeToString(b)
}

func generateCodeChallenge(verifier string) string {
    s256 := sha256.Sum256([]byte(verifier))
    return base64.RawURLEncoding.EncodeToString(s256[:])
}