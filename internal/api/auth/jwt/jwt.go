package jwt

import (
	"fmt"
	"os"
	"strings"
	"terrapak/internal/api/auth/roles"
	"terrapak/internal/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)



func DecodeJWT(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	secret := []byte(os.Getenv(config.ENV_TP_SECRET))
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	return claims, err

}

func GenerateJWT(user_id string, role roles.UserRoles) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = user_id
	claims["scope"] = role.String()
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv(config.ENV_TP_SECRET))
	return token.SignedString(secret)
}

func ParseToken(c *gin.Context) (string, error){
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(401)
		return "", fmt.Errorf("No Authorization header provided")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		c.AbortWithStatus(401)
		return "", fmt.Errorf("Invalid token")
	}

	return splitToken[1], nil
}