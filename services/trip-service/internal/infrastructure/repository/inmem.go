package repository

import (
	"context"

	"github.com/Minhajxdd/Synch/services/trip-service/internal/domain"
)

type inmemRepository struct {
	trip      map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		trip:      make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trip[trip.ID.Hex()] = trip
	return trip, nil
}
