package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = ":8000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
		fmt.Println("endpoint hello world hit")
	})

	fmt.Printf("starting server at %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
