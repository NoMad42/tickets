package v1

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (a apiServer) CreateBooking(w http.ResponseWriter, r *http.Request) {
	newBookingUuid, err := uuid.New().MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	w.Write(newBookingUuid)
}

func (a apiServer) GetBookingList(w http.ResponseWriter, r *http.Request) {
	// implement me
}
