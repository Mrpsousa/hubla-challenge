package internal

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Name  string `json:"name"`
	Idade int    `json:"idade"`
}

type BaseHandler struct {
	Base BaseStruct
}

func NewBaseHandler(base BaseStruct) *BaseHandler {
	return &BaseHandler{
		Base: base,
	}
}

func (b *BaseHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	get, _ := b.Base.DoGet()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(get)
}

func (b *BaseHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	data := &Data{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	post, _ := b.Base.DoPost(*data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
