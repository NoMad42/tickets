package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (a apiServer) GetFlightsList(w http.ResponseWriter, r *http.Request) {
	f, _ := a.flightsService.GetFlightsList(context.Background())

	response, err := json.Marshal(f)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}
