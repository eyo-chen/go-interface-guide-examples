package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eyo-chen/a-straignforward-guid-for-go-interface/example04-decoupling/without-interface/service"
)

// Handler is a handler to handle the request and response
type Handler struct {
	service service.MySQLService
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract the id from the URL query parameters
	id := r.URL.Query().Get("id")

	// Validate the id
	if !isValidID(id) {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get the user name from the service
	userName, err := h.service.GetUserName(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving user name: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userName)
}

func isValidID(id string) bool {
	return id != ""
}
