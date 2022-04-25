package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Users API")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
