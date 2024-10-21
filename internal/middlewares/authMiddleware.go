package middlewares

import (
	"go-reserve/internal/database"
	"go-reserve/internal/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizeUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			log.Printf("Authorization header is required |%v|", token)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		dbService := c.MustGet("db").(database.Service)
		db := dbService.GetDB()

		user, err := models.ValidateToken(db, token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userId", user.ID)
		c.Next()
	}
}
