package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Go trabalha com registro(Strusct) e não classe
//Struct modela os dados da entidade
//Sempre declarar maiusculo a inicial da propriedade para ter acesso fora do Struct ou fonte criado
type Contato struct {
	ID       int    `json:id`
	Nome     string `json:nome`
	Telefone string `json:telefone`
	Email    string `json:email`
}

//Criar arquivo main.go
//go mod init nome-projeto
//Gorilla/mux para trabalhar com http dentro do Golang
//go get github.com/gorilla/mux

//Vai representar o banco de dados
var contatos []Contato

func main() {

	//Criando contatos dentro do Array
	contatos = append(contatos, Contato{ID: 1, Nome: "John", Telefone: "(18)99999-9999", Email: "johnbuenodev@gmail.com"},
		Contato{ID: 2, Nome: "Marco Antonio", Telefone: "(18)88888-8888", Email: "toinhoborracheiro@gmail.com"},
		Contato{ID: 3, Nome: "Larinha", Telefone: "(18)77777-7777", Email: "magali@gmail.com"})

	//Formas de declarar var rota *mux.Router = mux.NewRouter()
	//Abaixo segue a inferencia de tipo
	rota := mux.NewRouter()

	rota.HandleFunc("/contatos", GetContatos).Methods("GET")
	rota.HandleFunc("/contatos/{id}", GetContatoById).Methods("GET")

	//Subir o servidor
	//Passando a porta que vai executar e as rotas declaradas
	//http.ListenAndServe(":8000", rota)

	log.Fatal(http.ListenAndServe(":8000", rota))
}

func GetContatos(response http.ResponseWriter, request *http.Request) {

	//converter dados para o Response que será enviado ao cliente solicitante
	//Enconder no response , converter/mapear o objeto contatos em json
	json.NewEncoder(response).Encode(contatos)
}

//PRECISA ADICIONAR VERIFICAÇÃO POR VALOR NULLO VAZIO NA REQUEST E RESPONSE PARA VALORES NÃO ENCONTRADOS
func GetContatoById(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	//a função strconv.Atoi retorna dois valores o id que foi solicitado da extração do params e o erro

	idContato, err := strconv.Atoi(params["id"])

	if err == nil {
		//Não vou usar a variavel i na programação então utilizo o _ underline para evitar um erro ele interpretar que não é um erro na compilação
		for _, contato := range contatos {
			if contato.ID == idContato {
				json.NewEncoder(response).Encode(contato)
			}
		}
	}

}
