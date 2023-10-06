package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"trabalho1/controller"
)

func main() {
	//cria as rotas
	r := chi.NewRouter()
	//People
	r.Post("/api/person/create", controller.CriarPessoa)          //Post Cria.
	r.Put("/api/person/{id}/update", controller.AtualizarPessoa)  //Put Modifica.
	r.Delete("/api/person/{id}/delete", controller.DeletarPessoa) //Delete Deleta
	r.Get("/api/person/{id}/search", controller.BuscarPessoa)     //Get procura/busca
	r.Get("/api/person/searchall", controller.BuscarPessoaTodos)

	//Products
	r.Post("/api/products/create", controller.CriarProduto)
	r.Put("/api/products/{id}/update", controller.AtualizarProduto)
	r.Delete("/api/products/{id}/delete", controller.DeletarProduto)
	r.Get("/api/products/{id}/search", controller.BuscarProduto)
	r.Get("/api/products/searchall", controller.BuscarProdutoTodos)

	//Orders
	r.Post("/api/orderpay/create", controller.CriarOrdem)
	r.Get("/api/orderpay/list-open-orders", controller.ListarOrdemAberta)
	r.Get("/api/ordepay/list-closed-orders", controller.ListarOrdemFechada)
	r.Put("/api/orderpay/{numberOrder}/insert-products", controller.AdicionarProdutoOrdem)
	r.Delete("/api/orderpay/{numberOrder}/{productId}/delete-products", controller.DeletarProdutoOrdem)

	//escuta a porta
	http.ListenAndServe(":8080", r)
}
