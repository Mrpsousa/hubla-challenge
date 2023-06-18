package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/database"
)

// 3. Exibir a lista de todas as transações de produtos importadas
// 4. Exibir o saldo final do produtor
// 5. Exibir o saldo final de um afiliado
//valor das transações em centavos /

type ListHandler struct {
	ListDB database.TransactionInterface
}

func NewListHandler(db database.TransactionInterface) *ListHandler {
	return &ListHandler{
		ListDB: db,
	}
}

func (t *ListHandler) ListProductorsBalance(w http.ResponseWriter, r *http.Request) {
	var producers []entity.DtoQueryResult

	producers, err := t.ListDB.ListProductorBalance()
	for idx := 0; idx < len(producers); idx++ {
		if producers[idx].TValue == 0 {
			producers = append(producers[:idx], producers[idx+1:]...)
			idx--
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(producers)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(producers)

}

func (t *ListHandler) ListAssociatesBalance(w http.ResponseWriter, r *http.Request) {
	var associates []entity.DtoQueryResult

	associates, err := t.ListDB.ListAssociateBalance()
	for idx := 0; idx < len(associates); idx++ {
		if associates[idx].TValue == 0 {
			associates = append(associates[:idx], associates[idx+1:]...)
			idx--
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(associates)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(associates)

}
