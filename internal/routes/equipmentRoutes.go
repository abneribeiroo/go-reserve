package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-reserve/internal/models"
)

func CreateEquipment(c *gin.Context, db *sql.DB) {
	var equipment models.Equipment
	if err := c.BindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := equipment.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create equipment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Equipment created"})
}
