package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	Id                primitive.ObjectID
	UserId            string
	PackageSlug       string
	TotalPriceInCents float64
}
