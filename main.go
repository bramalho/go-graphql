package main

import (
	"go-graphql/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Info).Methods("GET")
	r.HandleFunc("/health", handlers.Health).Methods("GET")
	r.HandleFunc("/graphql", handlers.GraphQL)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "80"
	}

	log.Println("Listning at http://localhost:" + port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
