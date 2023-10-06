package controller

import (
	"encoding/json"
	_ "io"
	"net/http"
	_ "trabalho1/models"
	"trabalho1/service"
)

func ListarOrdemAberta(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	resultOrder, err := service.ListOrderOpen()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(resultOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
