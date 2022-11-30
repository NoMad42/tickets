package v1

import (
	"context"
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
	t, _ := a.transactionsService.GetTransactionsList(context.Background())
	a.writeSuccessResponse(t, w)
}
