package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/database"
)

type ListHandler struct {
	ListDB database.TransactionInterface
}

func NewListHandler(db database.TransactionInterface) *ListHandler {
	return &ListHandler{
		ListDB: db,
	}
}

func (t *ListHandler) ListProductorsBalance(w http.ResponseWriter, r *http.Request) {
	var producers = make([]entity.DtoSellers, 0)

	producers, err := t.ListDB.GetProductorBalance()
	for idx := 0; idx < len(producers); idx++ {
		if producers[idx].TValue == 0 {
			producers = append(producers[:idx], producers[idx+1:]...)
			idx--
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(producers)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(producers)

}

func (t *ListHandler) ListAssociatesBalance(w http.ResponseWriter, r *http.Request) {
	var associates = make([]entity.DtoSellers, 0)

	associates, err := t.ListDB.GetAssociateBalance()
	for idx := 0; idx < len(associates); idx++ {
		if associates[idx].TValue == 0 {
			associates = append(associates[:idx], associates[idx+1:]...)
			idx--
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(associates)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(associates)

}
