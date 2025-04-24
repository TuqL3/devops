package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"go-devops/api/handlers"
	"go-devops/api/middlewares"
)

// SetupRouter thiết lập tất cả các routes
func SetupRouter(db *gorm.DB, logger *logrus.Logger) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middlewares.Logger(logger))

	// Metrics endpoint cho Prometheus
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// User handlers
		userHandler := handlers.NewUserHandler(db)
		users := api.Group("/users")
		{
			users.GET("/", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("/", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return router
}
