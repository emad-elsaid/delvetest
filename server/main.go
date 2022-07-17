package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Will start a server at 8090")
	http.ListenAndServe(":8090", nil)
}
