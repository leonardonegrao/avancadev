package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Printf("4º microsserviço.")
}
