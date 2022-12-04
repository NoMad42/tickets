package v1

import (
	"context"
	"net/http"
)

func (a apiServer) GetSeatsList(w http.ResponseWriter, r *http.Request) {
	s, _ := a.seatsService.GetSeatsList(context.Background())
	a.writeSuccessResponse(s, w)
}
