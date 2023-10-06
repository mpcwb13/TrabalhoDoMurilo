package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
	"trabalho1/models"
	"trabalho1/service"
)

func CriarProduto(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(readErr.Error()))
		return
	}

	var product models.Products
	err := json.Unmarshal(body, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	createProduct, err := service.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(createProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func AtualizarProduto(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(readErr.Error()))
		return
	}

	var products models.Products
	err := json.Unmarshal(body, &products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	//Transformando o id de string em int
	key := chi.URLParam(r, "id")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	//Chamando a função da regra de negocio.
	updateProduct, err := service.UpdateProduct(intKey, products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(updateProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func BuscarProduto(w http.ResponseWriter, r *http.Request) {

	key := chi.URLParam(r, "id")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	searchProduct, err := service.SearchProduct(intKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(searchProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func BuscarProdutoTodos(w http.ResponseWriter, r *http.Request) {

	searchAllProducts, err := service.SearchAllProducts()

	result, err := json.Marshal(searchAllProducts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {

	key := chi.URLParam(r, "id")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	service.DeletedProduct(intKey, models.Products{})

	w.WriteHeader(http.StatusOK)

}
