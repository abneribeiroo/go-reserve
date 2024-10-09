package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-reserve/internal/models"
)

func CreateReservation(c *gin.Context, db *sql.DB) {
	var reservation models.Reservation
	if err := c.BindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("INSERT INTO reservations (user_id, equipment_id, start_time, end_time, status) VALUES ($1, $2, $3, $4, $5)", reservation.UserID, reservation.EquipmentID, reservation.StartTime, reservation.EndTime, reservation.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reservation created"})
}
