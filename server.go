package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":4000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Fatal(http.ListenAndServe(PORT, nil))
}
