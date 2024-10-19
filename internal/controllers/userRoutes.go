package controllers

import (
	"log"
	"net/http"
	"strconv"

	"go-reserve/internal/database"
	"go-reserve/internal/models"

	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	
	dbService := c.MustGet("db").(database.Service) 
	db := dbService.GetDB()

	
	existingUser, _ := models.GetUserByEmail(db, user.Email)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}


	if err := user.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func GetAllUsers(c *gin.Context) {
	// Recupera a conexão do banco de dados do contexto
	dbService := c.MustGet("db").(database.Service) // Aqui você pega a instância da interface
	db := dbService.GetDB()

	users, err := models.GetAllUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Login autentica um usuário
func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Recupera a conexão do banco de dados do contexto
	dbService := c.MustGet("db").(database.Service) // Aqui você pega a instância da interface
	db := dbService.GetDB()

	user, err := models.GetUserByEmail(db, loginData.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !user.ComparePassword(loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetUserById retorna um usuário pelo ID
func GetUserById(c *gin.Context) {
	userIdParam := c.Param("userId")
    if userIdParam == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
        return
    }

    userId, err := strconv.Atoi(userIdParam)
    log.Println(userId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        log.Println(err)
        return
    }

	// Recupera a conexão do banco de dados do contexto
	dbService := c.MustGet("db").(database.Service) // Aqui você pega a instância da interface
	db := dbService.GetDB()

	user, err := models.GetUserById(db, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}


func UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user.ID = userId

	// Recupera a conexão do banco de dados do contexto
	dbService := c.MustGet("db").(database.Service) // Aqui você pega a instância da interface
	db := dbService.GetDB()

	if err := user.Update(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deleta um usuário pelo ID
func DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Recupera a conexão do banco de dados do contexto
	dbService := c.MustGet("db").(database.Service) // Aqui você pega a instância da interface
	db := dbService.GetDB()

	if err := models.DeleteUser(db, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted successfully"})
}
