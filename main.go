package main

import (
	"errors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv, cleanupFunc := di()
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-shutdown
		log.Println("Graceful Shutdown")
		cleanupFunc()
		log.Println("Cleared Resources")
		os.Exit(0)
	}()

}
