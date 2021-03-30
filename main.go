package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//implementa a codificação e a decodificação
)

//essa função trará todas as solicitações ao nosso url raiz
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!") //essa informação será mostrada na tela
	fmt.Println("Endpoint Hit: homePage")
}

//essa função combinará o acerto do caminho do URL com uma função definida
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil)) //a será inicializada na porta 10000
}

//função que iniciará a API
func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello1", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}

type Article struct { //struct com 3 propriedades para representar todos os artigos do site
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var Articles []Article //array global de artigos, será preenchida na função principal para simular o banco de dados

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

//parei na recuperação de arquivos
