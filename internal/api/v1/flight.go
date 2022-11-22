package v1

import (
	"context"
	"net/http"
)

func (a apiServer) GetFlightsList(w http.ResponseWriter, r *http.Request) {
	f, _ := a.flightsService.GetFlightsList(context.Background())
	a.writeSuccessResponse(f, w)
}
