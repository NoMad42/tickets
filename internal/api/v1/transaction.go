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
	w.Write(newTransactionUuid)
}

func (a apiServer) GetTransactionsList(w http.ResponseWriter, r *http.Request) {
	// implement me
}
