package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darwin808/greatest-todo/cmd/greatest-todo/db"
	"github.com/darwin808/greatest-todo/cmd/greatest-todo/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/", h.AddBook).Methods(http.MethodPost)
	fmt.Println("Established a successful connection!")

	log.Println("API is running!")
	http.ListenAndServe(":3000", router)
}
