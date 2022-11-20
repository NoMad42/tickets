package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (a apiServer) GetAirportsList(w http.ResponseWriter, r *http.Request) {
	ap, _ := a.airportsService.GetAirportsList(context.Background())

	response, err := json.Marshal(ap)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}
