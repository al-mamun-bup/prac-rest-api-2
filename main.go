package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World!"},
		Article{Title: "Test Title2", Desc: "Test Description2", Content: "Hello World!2"},
	}
	fmt.Println("Endpoint Hit! : All Articles Endpoint ")
	json.NewEncoder(w).Encode(articles)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage End Point Hit!")
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//myRouter.HandleFunc()
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
func main() {
	handleRequests()
}
