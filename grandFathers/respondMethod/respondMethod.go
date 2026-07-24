package respondMethod

import (
	"encoding/json"
	"net/http"
)

func SendRespond(w http.ResponseWriter, r *http.Request, respond any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respond)
}
