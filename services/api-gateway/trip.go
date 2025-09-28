package main

import (
	"encoding/json"
	"net/http"

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

	data := contracts.APIResponse{Data: "ok"}

	writeJSON(w, http.StatusCreated, data)
}
