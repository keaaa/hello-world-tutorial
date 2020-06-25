package handlers

import (
	"fmt"
	"net/http"
)

func GetHelloWorldMsg() string {
	return "Hello world!"
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, GetHelloWorldMsg())
	fmt.Println("endpoint hello world hit")
}
