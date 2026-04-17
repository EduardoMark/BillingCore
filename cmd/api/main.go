package main

import (
	"context"
	"net/http"

	"github.com/EduardoMark/BillingCore/internal/account"
	"github.com/EduardoMark/BillingCore/internal/billing/customer"
	"github.com/EduardoMark/BillingCore/internal/billing/plans"
	"github.com/EduardoMark/BillingCore/internal/billing/subscription"
	"github.com/EduardoMark/BillingCore/internal/cache"
	"github.com/EduardoMark/BillingCore/internal/infra/database"
	"github.com/EduardoMark/BillingCore/internal/infra/rabbitmq"
	"github.com/EduardoMark/BillingCore/internal/infra/rediscache"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	db, err := database.New()
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

	BindingRoutes(router, db)

	// Start the server
	router.Run(":8080")
}

func BindingRoutes(router *gin.Engine, db *gorm.DB) {
	ctx := context.Background()

	api := router.Group("/api/v1")
	accGroup := api.Group("/accounts")

	// redis client
	redisClient := rediscache.NewRedisCache()
	redisCacheClient := cache.NewRedisCache(redisClient)

	// producer
	producer := rabbitmq.NewProducer()

	// account routes
	accRepository := account.NewRepository(db)
	accService := account.NewService(accRepository)
	accHandler := account.NewHandler(accService)
	accHandler.RegisterRoutes(accGroup)

	// plan routes
	planRepository := plans.NewRepository(db)
	planService := plans.NewService(planRepository)
	planHandler := plans.NewHandler(planService)
	planHandler.RegisterRoutes(accGroup)

	// customer routes
	custRepository := customer.NewRepository(db)
	custService := customer.NewService(custRepository)
	custHandler := customer.NewHandler(custService)
	custHandler.RegisterRoutes(accGroup)

	// subscription routes
	subRepo := subscription.NewRepository(db)
	subService := subscription.NewService(subRepo, producer)
	subConsumer := subscription.NewConsumer(subService)
	subHandler := subscription.NewHandler(subService, redisCacheClient)
	subHandler.RegisterRoutes(accGroup)

	go func() {
		if err := rabbitmq.Consume(ctx, "subscription.created", subConsumer.Handle); err != nil {
			zap.L().Error("consumer error", zap.Error(err))
		}
	}()
}
