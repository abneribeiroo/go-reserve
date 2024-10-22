package controllers

import (
	"net/http"
	"strconv"
	"time"

	"go-reserve/internal/database"
	"go-reserve/internal/models"

	"github.com/gin-gonic/gin"
)

// Criação de uma nova reserva
func CreateReservation(c *gin.Context) {
	var input struct {
		EquipmentID int       `json:"equipment_id"`
		StartTime   time.Time `json:"start_time"`
		EndTime     time.Time `json:"end_time"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, _ := strconv.Atoi(c.Param("userId"))

	reservation := models.NewReservation(userID, input.EquipmentID, input.StartTime, input.EndTime)

	dbService := c.MustGet("db").(database.Service)
	db := dbService.GetDB()

	if err := reservation.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reservation created successfully", "reservation": reservation})
}

// Aprovar uma reserva
func ApproveReservation(c *gin.Context) {
	reservationID, _ := strconv.Atoi(c.Param("reservationId"))

	dbService := c.MustGet("db").(database.Service)
	db := dbService.GetDB()

	var reservation models.Reservation
	if err := db.QueryRow("SELECT * FROM reservations WHERE id = $1", reservationID).Scan(&reservation.ID, &reservation.UserID, &reservation.EquipmentID, &reservation.StartTime, &reservation.EndTime, &reservation.Status, &reservation.CreatedAt); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	if err := reservation.Approve(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation approved successfully"})
}

// Rejeitar uma reserva
func RejectReservation(c *gin.Context) {
	reservationID, _ := strconv.Atoi(c.Param("reservationId"))

	dbService := c.MustGet("db").(database.Service)
	db := dbService.GetDB()

	var reservation models.Reservation
	if err := db.QueryRow("SELECT * FROM reservations WHERE id = $1", reservationID).Scan(&reservation.ID, &reservation.UserID, &reservation.EquipmentID, &reservation.StartTime, &reservation.EndTime, &reservation.Status, &reservation.CreatedAt); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	if err := reservation.Reject(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation rejected successfully"})
}
