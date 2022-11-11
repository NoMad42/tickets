package v1

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (a apiServer) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	newTransactionUuid, err := uuid.New().MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(newTransactionUuid)
	if err != nil {
		log.Fatal(err)
	}
}

func (a apiServer) GetTransactionsList(w http.ResponseWriter, r *http.Request) {
	// implement me
}
