package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/Bobby-P-dev/search-enggine-go/src/models"
)

func WriteJSON(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(payload)
}

func SuccesResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	_ = WriteJSON(w, statusCode, models.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	_ = WriteJSON(w, statusCode, models.APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
	})
}
