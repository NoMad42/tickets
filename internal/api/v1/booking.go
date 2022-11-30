package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (a apiServer) CreateBooking(w http.ResponseWriter, r *http.Request) {
	newBookingUuid, err := uuid.New().MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(newBookingUuid)
	if err != nil {
		log.Fatal(err)
	}
}

func (a apiServer) GetBookingList(w http.ResponseWriter, r *http.Request) {
	bl, _ := a.bookingService.GetBookingList(context.Background())
	a.writeSuccessResponse(bl, w)
}
