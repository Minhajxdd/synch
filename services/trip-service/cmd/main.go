package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Minhajxdd/Synch/services/trip-service/internal/infrastructure/repository"
	"github.com/Minhajxdd/Synch/services/trip-service/internal/service"
	grpcserver "google.golang.org/grpc"
)

var GrpcAddr = ":9093"

func main() {
	inMemRepo := repository.NewInmemRepository()
	svc := service.NewService(inMemRepo)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}

	grpcServer := grpcserver.NewServer()

	log.Println("starting grpc server on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("Failed to start grpc server")
			cancel()
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}
