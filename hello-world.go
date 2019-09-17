package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Sigilita/GoTest/types"
)

var articles types.Articles

func main() {
	articles = types.Articles{
		types.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		types.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"}}
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: ReturnAllArticles")
	json.NewEncoder(w).Encode(articles)
}
