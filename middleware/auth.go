package middleware

import (
	"context"
	"net/http"
	"solana-balance-api/models"
	"solana-balance-api/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing API Key"})
			return
		}
		collection := utils.MongoClient.Database("solana").Collection("apikeys")
		var result models.APIKey
		err := collection.FindOne(context.Background(), map[string]string{"key": apiKey}).Decode(&result)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			return
		}
		c.Next()
	}
}