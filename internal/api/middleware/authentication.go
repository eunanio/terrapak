package middleware

import (
	"terrapak/internal/api/auth"
	"terrapak/internal/api/auth/roles"

	"github.com/gin-gonic/gin"
)

func HasAuthenticatedRole(role roles.UserRoles) gin.HandlerFunc {
	return func(c *gin.Context) {
		authProvider := auth.GetAuthProvider()
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		authProvider.Authenticate(token)


	}
}