package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	Id                primitive.ObjectID
	UserId            string
	PackageSlug       string
	TotalPriceInCents float64
}

type TripModel struct {
	ID       primitive.ObjectID
	UserId   string
	Status   string
	RideFare RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}
