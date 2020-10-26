package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/say-hello-world", SayHelloWorld)
	http.ListenAndServe(":9090", nil)
}

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
