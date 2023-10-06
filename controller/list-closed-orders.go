package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"trabalho1/models"
	"trabalho1/service"
)

func ListarOrdemFechada(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(readErr.Error()))
		return
	}

	var closedOrder models.OrderPay
	err := json.Unmarshal(body, &closedOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resultOrder, err := service.ListOrderClosed()
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
