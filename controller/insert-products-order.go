package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
	"trabalho1/service"
)

func AdicionarProdutoOrdem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(readErr.Error()))
		return
	}

	key := chi.URLParam(r, "numberOrder")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var productsId []int
	err = json.Unmarshal(body, &productsId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	insertProduct, err := service.InsertProducts(intKey, productsId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(insertProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
