package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-reserve/internal/models"
)

func CreateUser(c *gin.Context, db *sql.DB) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// Similarmente, outros métodos de login, listagem e exclusão podem ser implementados.
