package v1

import (
	"context"
	"encoding/json"
	"net/http"
)

func (a apiServer) GetAirportsList(w http.ResponseWriter, r *http.Request) {
	ap, _ := a.airportsService.GetAirportsList(context.Background())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ap)
}
