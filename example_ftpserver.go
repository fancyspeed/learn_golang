package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":18080", http.FileServer(http.Dir("E:\\go_workspace"))))
}