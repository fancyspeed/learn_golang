package main

import (
	"fmt"
	"net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        fmt.Fprint(w, "hello, post")
	} else {
		fmt.Fprint(w, "hello, get")
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":18082", nil)
}