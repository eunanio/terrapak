package middleware

import (
	"log/slog"
	"slices"
	"terrapak/internal/api/auth/jwt"
	"terrapak/internal/api/auth/roles"
	"terrapak/internal/db/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HasAuthenticatedRole(roles ...roles.UserRoles) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		authHeader := c.GetHeader("Authorization")
		us := services.UserService{}
		if authHeader == "" {
			slog.Debug("No Authorization header provided")
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		token, err := jwt.ParseToken(c); if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		claims,err := jwt.DecodeJWT(token); if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		user_id, err := uuid.Parse(claims["id"].(string)); if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		user := us.Find(user_id)
		if user == nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		if !slices.Contains(roles, user.Role) {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		c.Next()
	}
}