package main

import (
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /health", handlerReadiness)

	serverStruct := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(serverStruct.ListenAndServe())
}
