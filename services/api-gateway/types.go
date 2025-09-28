package main

import "github.com/Minhajxdd/Synch/shared/types"

type previewTripRequest struct {
	UserId      string           `json:"userId"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
