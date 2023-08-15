package main

import (
	"context"
	"dbgo/wallet"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func ProvideServer(handler *gin.Engine) (*http.Server, func()) {
	server := &http.Server{
		Handler: handler,
		Addr:    ":8080",
	}
	return server, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("can not shutdown")
		}
	}
}

func RegisterWallet(r *gin.Engine, handler *wallet.Handler) {
	r.POST("/wallets", handler.SaveWallet)
	r.GET("/wallets/:id", handler.GetWalletByID)
	r.GET("/wallets/:id/balance", handler.GetBalanceByID)
	r.POST("/wallets/:id/deposit", handler.DepositByID)
	r.POST("/wallets/:id/withdraw", handler.WithdrawByID)
}
