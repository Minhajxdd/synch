package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	Id                primitive.ObjectID
	UserID            string
	PackageSlug       string
	TotalPriceInCents float64
}
