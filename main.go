package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /health", handlerHealth)

	serverStruct := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("server starting on: localhost:8080")
	log.Fatal(serverStruct.ListenAndServe())
}
