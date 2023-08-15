package main

import (
	"dbgo/db"
	"dbgo/wallet"
	"github.com/gin-gonic/gin"
	"net/http"
)

func di() (*http.Server, func()) {
	conn, cleanupDBConnFunc := db.ProvideDBCon()
	walletRepo := wallet.ProvideRepo(conn)
	walletHandler := wallet.ProvideHandler(walletRepo)
	ginEngine := gin.Default()
	RegisterWallet(ginEngine, walletHandler)
	srv, cleanupServerFunc := ProvideServer(ginEngine)

	return srv, func() {
		cleanupDBConnFunc()
		cleanupServerFunc()
	}

}
