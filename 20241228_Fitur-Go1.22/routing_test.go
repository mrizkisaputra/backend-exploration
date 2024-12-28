package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRoutingGo_1_22(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	mux.HandleFunc("GET /product/{id_product}", func(writer http.ResponseWriter, request *http.Request) {
		id := request.PathValue("id_product")
		fmt.Fprintf(writer, "Your id is %s", id)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		t.Fatal(err)
	}
}
