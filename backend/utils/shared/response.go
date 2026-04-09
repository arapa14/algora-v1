package shared

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIMessage struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondJSON(w http.ResponseWriter, status int, statusMsg string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(APIMessage{
		Status:  statusMsg,
		Message: message,
		Data:    data,
	})

	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RespondError(w http.ResponseWriter, status int, message string) {
	log.Printf("❌ [Error %d]: %s", status, message)
	RespondJSON(w, status, "Error", message, nil)
}

func RespondSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	log.Printf("✅ [Success %d]: %s", status, message)
	RespondJSON(w, status, "Success", message, data)
}
