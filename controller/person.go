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

func CriarPessoa(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()                // Aqui estou fechando o Body, mas por estar usando defer, vai ser fechado no final da função
	body, readErr := io.ReadAll(r.Body) // Aqui estou lendo tudo que está Body da requisição e armazenando na variável body e vendo se o ReadAll teve erro ou nao.
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(readErr.Error()))
		return
	}

	var person models.Person
	//Chamando função, lendo o body, e armazenando o resultado dentro de &person e guardando em err o resultado da chamada da função
	err := json.Unmarshal(body, &person) //Desserializando o que vem do body através da requisiçao
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Escrevendo o status da resposta para o http
		w.Write([]byte(err.Error()))         // Estou descrevendo o erro para o http
		return
	}

	createPerson, err := service.CreatePerson(person) //Aqui está usando a função "CreatePerson"(Que vem de service) e criando a pessoa
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//Chamando a função para serializar ''person''
	result, err := json.Marshal(createPerson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Escrevendo o status da resposta para o http
		w.Write([]byte(err.Error()))         // Estou descrevendo o erro para o http
		return
	}

	w.WriteHeader(http.StatusOK) //Enviando para o http o status ok
	w.Write(result)              //Enviando uma resposta para a request escrita em json
}

func AtualizarPessoa(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()                // Aqui estou fechando o Body, mas por estar usando defer, vai ser fechado no final da função
	body, readErr := io.ReadAll(r.Body) // Aqui estou lendo tudo que está no Body da requisição e armazenando na variável body e vendo se o ReadAll teve erro ou nao.
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(readErr.Error()))
		return
	}

	var people models.Person
	//Chamando função, lendo o body, e armazenando o resultado dentro de &person e guardando em err o resultado da chamada da função
	err := json.Unmarshal(body, &people)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Escrevendo o status da resposta para o http
		w.Write([]byte(err.Error()))         // Estou descrevendo o erro para o http
		return
	}

	key := chi.URLParam(r, "id")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}

	updatePerson, err := service.UpdatePerson(intKey, people)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//Chamando a função para serializar ''person''
	result, err := json.Marshal(updatePerson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Escrevendo o status da resposta para o http
		w.Write([]byte(err.Error()))         // Estou descrevendo o erro para o http
		return
	}

	w.WriteHeader(http.StatusOK) //Enviando para o http o status ok
	w.Write(result)              //Enviando uma resposta para a request escrita em json

}

func DeletarPessoa(w http.ResponseWriter, r *http.Request) {

	key := chi.URLParam(r, "id")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}

	service.DeletedPeople(intKey, models.Person{})

	w.WriteHeader(http.StatusOK) //Enviando para o http o status ok

}

func BuscarPessoa(w http.ResponseWriter, r *http.Request) {

	key := chi.URLParam(r, "id")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	SearchPeople, err := service.SearchPeople(intKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(SearchPeople)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Escrevendo o status da resposta para o http
		w.Write([]byte(err.Error()))         // Estou descrevendo o erro para o http
		return
	}

	w.WriteHeader(http.StatusOK) //Enviando para o http o status ok
	w.Write(result)

}

func BuscarPessoaTodos(w http.ResponseWriter, r *http.Request) {

	SearchAllPeople, err := service.SearchAllPeople()

	result, err := json.Marshal(SearchAllPeople)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Escrevendo o status da resposta para o http
		w.Write([]byte(err.Error()))         // Estou descrevendo o erro para o http
		return
	}

	w.WriteHeader(http.StatusOK) //Enviando para o http o status ok
	w.Write(result)

}
