package main

import (
	"time"

	"github.com/Minhajxdd/Synch/services/trip-service/internal/infrastructure/repository"
	"github.com/Minhajxdd/Synch/services/trip-service/internal/service"
)

func main() {
	inMemRepo := repository.NewInmemRepository()
	_ = service.NewService(inMemRepo)

	for {
		time.Sleep(time.Second)
	}
}
