package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williamkoller/7.1-crud/server"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)

	fmt.Println("listen port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
