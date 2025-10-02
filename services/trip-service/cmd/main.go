package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	h "github.com/Minhajxdd/Synch/services/trip-service/internal/infrastructure/http"
	"github.com/Minhajxdd/Synch/services/trip-service/internal/infrastructure/repository"
	"github.com/Minhajxdd/Synch/services/trip-service/internal/service"
)

func main() {
	inMemRepo := repository.NewInmemRepository()
	svc := service.NewService(inMemRepo)
	mux := http.NewServeMux()

	httphandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httphandler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		fmt.Printf("Failed to start server : %v", err)

	case sig := <-shutdown:
		fmt.Printf("Server is shutting down due to : %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Unable to shutdown gracefully : %v", err)
			server.Close()
		}
	}
}
