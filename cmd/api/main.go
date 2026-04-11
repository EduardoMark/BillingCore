package main

import (
	"net/http"

	"github.com/EduardoMark/BillingCore/internal/infra/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env file", zap.Error(err))
	}

	// Initialize database connection
	_, err := database.New()
	if err != nil {
		logger.Fatal("Error connecting to database", zap.Error(err))
	}

	if err := database.Migrate(); err != nil {
		logger.Fatal("Error on migrate database", zap.Error(err))
	}

	// Initialize Gin router
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.Run(":8080")
}
