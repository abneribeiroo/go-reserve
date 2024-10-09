package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-reserve/internal/controllers"
	
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(func (c *gin.Context) { 
		c.Set("db", s.db) 
		c.Next() 
	})

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	// Prefixo API v1
	api := r.Group("/api/v1")
	{
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/register", controllers.CreateUser)
			authRoutes.POST("/login", controllers.Login)
		}

		userRoutes := api.Group("/users")
		{
			userRoutes.GET("/", controllers.GetAllUsers)
			userRoutes.GET("/:userId", controllers.GetUserById)
			userRoutes.PUT("/:userId", controllers.UpdateUser)
			userRoutes.DELETE("/:userId", controllers.DeleteUser)
		}

		// userReservationRoutes := userRoutes.Group("/:userId/reservations")
		// {
		// 	userReservationRoutes.POST("/", controllers.CreateReservation)             
		// 	userReservationRoutes.GET("/", controllers.GetAllReservationsForUser)        
		// 	userReservationRoutes.GET("/:reservationId", controllers.GetReservationById)  
		// 	userReservationRoutes.PUT("/:reservationId", controllers.UpdateReservation)  
		// 	userReservationRoutes.DELETE("/:reservationId", controllers.DeleteReservation) 
		// }

		equipmentRoutes := api.Group("/equipments")
		{
			equipmentRoutes.POST("/", controllers.CreateEquipment)
			equipmentRoutes.GET("/", controllers.GetAllEquipment)
			equipmentRoutes.GET("/:equipmentId", controllers.GetEquipmentById)
			equipmentRoutes.PUT("/:equipmentId", controllers.UpdateEquipment)
			equipmentRoutes.DELETE("/:equipmentId", controllers.DeleteEquipment)
		}
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
