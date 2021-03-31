package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" //implementa a codificação e a decodificação

	"github.com/gorilla/mux"
)

//essa função trará todas as solicitações ao nosso url raiz
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!") //essa informação será mostrada na tela
	fmt.Println("Endpoint Hit: homePage")
}

//essa função combinará o acerto do caminho do URL com uma função definida
func handleRequests() { //função que mapeará todas as chamadas
	myRouter := mux.NewRouter().StrictSlash(true)

	http.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", nil)) //a será inicializada na porta 10000
}

//função que iniciará a API
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}

type Article struct { //struct com 3 propriedades para representar todos os artigos do site
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var Articles []Article //array global de artigos, será preenchida na função principal para simular o banco de dados

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) //codifica a matriz de arquivos em uma string JSON, e em seguida, escreve como parte da resposta
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
