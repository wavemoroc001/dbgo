package main

import (
	"dbgo/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var cleanupDBConnFunc func()
	db.Conn, cleanupDBConnFunc = db.ProvideDBCon()
	router := gin.Default()
	Register(router)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdown
		cleanupDBConnFunc()
		os.Exit(0)
	}()

	log.Fatal(router.Run(":8080"))

}
