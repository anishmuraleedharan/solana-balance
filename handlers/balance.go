package handlers

import (
	"context"
	"net/http"

	//"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gin-gonic/gin"
)

type BalanceRequest struct {
	Wallets []string `json:"wallets"`
}

type BalanceResponse struct {
	Wallet  string `json:"wallet"`
	Balance uint64 `json:"balance"`
}

func GetBalanceHandler(c *gin.Context) {
	var req BalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	client := rpc.New(rpc.MainNetBeta_RPC)
	ctx := context.Background()

	var results []BalanceResponse
	for _, wallet := range req.Wallets {
		pubkey, err := solana.PublicKeyFromBase58(wallet)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet address: " + wallet})
			return
		}
		balanceResult, err := client.GetBalance(ctx, pubkey, rpc.CommitmentFinalized)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance"})
			return
		}

		results = append(results, BalanceResponse{
			Wallet:  wallet,
			Balance: balanceResult.Value,
		})
	}

	c.JSON(http.StatusOK, results)
}
