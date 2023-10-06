package controller

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"trabalho1/service"
)

func DeletarProdutoOrdem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	key := chi.URLParam(r, "numberOrder")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	key2 := chi.URLParam(r, "productId")
	intKey2, err := strconv.Atoi(key2)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	service.DeleteProductOrder(intKey, intKey2)

	w.WriteHeader(http.StatusOK)

}
