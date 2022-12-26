package v1

import (
	"context"
	"encoding/json"
	"homework/specs"
	"log"
	"net/http"
)

func (a apiServer) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tcr specs.TransactionCreateRequest
	err := json.NewDecoder(r.Body).Decode(&tcr)
	if err != nil {
		log.Println("JSON decode error: ", err)
	}
	defer r.Body.Close()

	u, err := a.transactionsService.CreateTransaction(context.Background(), tcr.BookingId)
	if err != nil {
		log.Println(err)
		return
	}

	a.writeSuccessResponse(u, w)
}

func (a apiServer) GetTransactionsList(w http.ResponseWriter, r *http.Request) {
	t, _ := a.transactionsService.GetTransactionsList(context.Background())
	a.writeSuccessResponse(t, w)
}
