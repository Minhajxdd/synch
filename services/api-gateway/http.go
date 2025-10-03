package main

import (
	"encoding/json"
	"log"
	"net/http"

	grpcclient "github.com/Minhajxdd/Synch/services/api-gateway/grpc_client"
	"github.com/Minhajxdd/Synch/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if reqBody.UserId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	tripService, err := grpcclient.NewTripServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	defer tripService.Close()

	tripPreview, err := tripService.Client.PreviewTrip(r.Context(), reqBody.toProto())
	if err != nil {
		http.Error(w, "Failed to preview trip", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	response := contracts.APIResponse{Data: tripPreview}

	writeJSON(w, http.StatusCreated, response)
}
