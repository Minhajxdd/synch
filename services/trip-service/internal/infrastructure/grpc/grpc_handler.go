package grpc

import (
	"context"
	"log"

	"github.com/Minhajxdd/Synch/services/trip-service/internal/domain"
	pb "github.com/Minhajxdd/Synch/shared/proto/trip"
	"github.com/Minhajxdd/Synch/shared/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	pb.UnimplementedTripServiceServer

	service domain.TripService
}

func NewGrpcHandler(server *grpc.Server, service domain.TripService) *grpcHandler {
	handler := &grpcHandler{
		service: service,
	}

	pb.RegisterTripServiceServer(server, handler)
	return handler
}

func (h *grpcHandler) PreviewTrip(ctx context.Context, req *pb.PreviewTripRequest) (*pb.PreviewTripResponse, error) {
	startLocation := req.GetStartLocation()
	endLocation := req.GetEndLocation()
	pickup := &types.Coordinate{
		Latitude:  startLocation.Latitude,
		Longitude: startLocation.Longitude,
	}
	destination := &types.Coordinate{
		Latitude:  endLocation.Latitude,
		Longitude: endLocation.Longitude,
	}

	t, err := h.service.GetRoute(ctx, pickup, destination)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to get route: %v", err)
	}

	return &pb.PreviewTripResponse{
		Route:     t.ToProto(),
		RideFares: []*pb.RideFare{},
	}, nil
}
