package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AbdullahBasir/financial-tracker/internal/handler"
	"github.com/AbdullahBasir/financial-tracker/internal/middleware"
)

func main() {
	cfg := handler.LoadConfig()

	router := http.NewServeMux()

	router.HandleFunc("GET /health", handler.HandlerHealth)

	router.HandleFunc("POST /auth/register", cfg.HandlerRegister)
	router.HandleFunc("POST /auth/login", cfg.HandlerLogin)

	router.Handle("POST /accounts", middleware.Auth(http.HandlerFunc(cfg.HandlerCreateAccount)))
	router.Handle("GET /accounts", middleware.Auth(http.HandlerFunc(cfg.HandlerGetAccounts)))
	router.Handle("GET /accounts/{id}", middleware.Auth(http.HandlerFunc(cfg.HandlerGetAccount)))
	router.Handle("PATCH /accounts/{id}", middleware.Auth(http.HandlerFunc(cfg.HandlerUpdateAccount)))
	router.Handle("DELETE /accounts/{id}", middleware.Auth(http.HandlerFunc(cfg.HandlerDeleteAccount)))

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
