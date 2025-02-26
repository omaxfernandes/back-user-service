package response

import (
	"encoding/json"
	"net/http"
)

// JSON envia uma resposta JSON
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
