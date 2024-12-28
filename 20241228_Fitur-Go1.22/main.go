package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	mux.HandleFunc("GET /product/{id_product}", func(writer http.ResponseWriter, request *http.Request) {
		id := request.PathValue("id_product")
		fmt.Fprintf(writer, "Your id product is %s", id)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}

}
