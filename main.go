package main

import (
	"log"
	"solana-balance-api/handlers"
	"solana-balance-api/middleware"
	"solana-balance-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginMiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memoryStore "github.com/ulule/limiter/v3/drivers/store/memory"
)

func main() {
	err := utils.InitMongo("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	rate, _ := limiter.NewRateFromFormatted("10-M")
	store := memoryStore.NewStore()
	rateLimiter := ginMiddleware.NewMiddleware(limiter.New(store, rate))

	router := gin.Default()
	router.Use(rateLimiter)
	router.Use(middleware.AuthMiddleware())

	//router.POST("/api/get-balance", handlers.GetBalanceHandler)
	router.POST("/api/get-balance", handlers.GetBalanceHandler)

	router.Run(":8080")
}
