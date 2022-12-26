package v1

import (
	"context"
	"encoding/json"
	"homework/specs"
	"log"
	"net/http"
)

func (a apiServer) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var bcr specs.BookingCreateRequest
	err := json.NewDecoder(r.Body).Decode(&bcr)
	if err != nil {
		log.Println("JSON decode error: ", err)
	}
	defer r.Body.Close()

	u, err := a.bookingService.CreateBooking(context.Background(), bcr.SeatId, bcr.UserProfileId)
	if err != nil {
		log.Println(err)
		return
	}

	a.writeSuccessResponse(u, w)
}

func (a apiServer) GetBookingList(w http.ResponseWriter, r *http.Request) {
	bl, _ := a.bookingService.GetBookingList(context.Background())
	a.writeSuccessResponse(bl, w)
}
