package v1

import (
	"context"
	"net/http"
)

func (a apiServer) GetAirportsList(w http.ResponseWriter, r *http.Request) {
	ap, _ := a.airportsService.GetAirportsList(context.Background())
	a.writeSuccessResponse(ap, w)
}
