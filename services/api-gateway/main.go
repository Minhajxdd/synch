package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Minhajxdd/Synch/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", Cors(handleTripPreview))
	mux.HandleFunc("/ws/riders", handleRidersWebsocket)
	mux.HandleFunc("/ws/drivers", handleDriversWebsocket)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	serverErrrors := make(chan error, 1)

	go func() {
		serverErrrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrrors:
		log.Printf("Failed to start server : %v", err)

	case sig := <-shutdown:
		log.Printf("Server is shutting down due to : %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Unable for gracefull shutdown : %v", err)
			server.Close()
		}
	}
}
