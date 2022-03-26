package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Products []Product

type productHandler struct {
	sync.Mutex
	products Products
}

func (ph *productHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ph.get(w, r)
	case "POST":
		ph.post(w, r)
	case "PUT", "PATCH":
		ph.put(w, r)
	case "DELETE":
		ph.delete(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid method")
	}
}

func (ph *productHandler) get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from GET")
}
func (ph *productHandler) post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from POST")

}
func (ph *productHandler) put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from PUT")
}
func (ph *productHandler) delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from DELETE")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	responseWithJson(w, code, map[string]string{"error": msg})
}

func responseWithJson(w http.ResponseWriter, code int, data interface{}) {
	response, err := json.Marshal(data)

	if err != nil {
		log.Fatalln("Erro ao converter para json")
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	PORT := ":4000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Fatal(http.ListenAndServe(PORT, nil))
}
