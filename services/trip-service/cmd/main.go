package main

import (
	"log"
	"net/http"

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

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
