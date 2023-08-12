package main

import (
	"dbgo/wallet"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/wallets", wallet.SaveWallet)
	r.GET("/wallets/:id", wallet.GetWalletByID)
	r.GET("/wallets/:id/balance", wallet.GetBalanceByID)
	r.POST("/wallets/:id/deposit", wallet.DepositByID)
}
