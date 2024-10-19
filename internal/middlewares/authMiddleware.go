package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o valor do cabeçalho "Authorization"
		tokenString := c.Request.Header.Get("Authorization")
		
		// Verifica se o token foi fornecido
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// Verifica se o token começa com "Bearer "
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token format"})
			c.Abort()
			return
		}

		// Remove o prefixo "Bearer " para obter apenas o token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Valida o token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica se o método de assinatura é esperado (neste caso, HMAC)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Retorna a chave secreta
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// Se o token for válido, prossegue
		c.Next()
	}
}
