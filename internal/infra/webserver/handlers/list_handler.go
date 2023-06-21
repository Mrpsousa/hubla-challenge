package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpsousa/api/internal/dto"
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

// ListProductorsBalance godoc
// @Summary      List productors balance
// @Description  get all productors balance
// @Tags         productors balance
// @Accept       json
// @Produce      json
// @Success      200       {array}   dto.DtoSellers
// @Failure      500       {object}  Error
// @Router       /producers [get]
// @Security ApiKeyAuth
func (t *ListHandler) ListProductorsBalance(w http.ResponseWriter, r *http.Request) {
	var producers = make([]dto.DtoSellers, 0)

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

// ListAssociatesBalance godoc
// @Summary      List associates balance
// @Description  get all associates balance
// @Tags         associates balance
// @Accept       json
// @Produce      json
// @Success      200       {array}   dto.DtoSellers
// @Failure      500       {object}  Error
// @Router       /associates [get]
// @Security ApiKeyAuth
func (t *ListHandler) ListAssociatesBalance(w http.ResponseWriter, r *http.Request) {
	var associates = make([]dto.DtoSellers, 0)

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

// ListForeignCourses godoc
// @Summary      List foreign courses
// @Description  get all foreign courses
// @Tags         foreign courses
// @Accept       json
// @Produce      json
// @Success      200       {array}   dto.DtoCourses
// @Failure      500       {object}  Error
// @Router       /courses/foreign [get]
// @Security ApiKeyAuth
func (t *ListHandler) ListForeignCourses(w http.ResponseWriter, r *http.Request) {
	var courses = make([]dto.DtoCourses, 0)

	courses, err := t.ListDB.GetForeignCourses()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(courses)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)

}
