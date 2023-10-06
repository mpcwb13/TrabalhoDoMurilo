package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"trabalho1/models"
	"trabalho1/service"
)

func CriarOrdem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(readErr.Error()))
		return
	}

	var orderPay models.OrderPay
	err := json.Unmarshal(body, &orderPay)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	createOrder, err := service.CreateOrder(orderPay)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(createOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
