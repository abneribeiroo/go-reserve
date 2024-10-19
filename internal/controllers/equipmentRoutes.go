package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"go-reserve/internal/database"
	"go-reserve/internal/models"
)


func CreateEquipment(c *gin.Context) {
	var equipment models.Equipment
	if err := c.BindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	
	dbService := c.MustGet("db").(database.Service) 
	db := dbService.GetDB()

	
	if err := equipment.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create equipment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Equipment created"})
}


func GetAllEquipment(c *gin.Context) {
	
	dbService := c.MustGet("db").(database.Service) 
	db := dbService.GetDB()

	equipmentList, err := models.GetAllEquipment(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve equipment"})
		return
	}

	c.JSON(http.StatusOK, equipmentList)
}

// GetEquipmentById retorna um equipamento pelo ID
func GetEquipmentById(c *gin.Context) {
	equipmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	// Recupera a conexão do banco de dados do contexto
	dbService := c.MustGet("db").(database.Service) // Aqui você pega a instância da interface
	db := dbService.GetDB()

	equipment, err := models.GetEquipmentById(db, equipmentId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipment not found"})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

// UpdateEquipment atualiza um equipamento pelo ID
func UpdateEquipment(c *gin.Context) {
	equipmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	var equipment models.Equipment
	if err := c.BindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	equipment.ID = equipmentId

	// Recupera a conexão do banco de dados do contexto
	dbService := c.MustGet("db").(database.Service) // Aqui você pega a instância da interface
	db := dbService.GetDB()

	// Atualiza o equipamento no banco de dados
	if err := equipment.Update(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update equipment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipment updated successfully"})
}

// DeleteEquipment deleta um equipamento pelo ID
func DeleteEquipment(c *gin.Context) {
	equipmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	// Recupera a conexão do banco de dados do contexto
	dbService := c.MustGet("db").(database.Service) // Aqui você pega a instância da interface
	db := dbService.GetDB()

	// Deleta o equipamento do banco de dados
	if err := models.DeleteEquipment(db, equipmentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete equipment"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Equipment deleted successfully"})
}
