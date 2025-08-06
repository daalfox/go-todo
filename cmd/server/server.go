package main

import (
	"log"
	"net/http"

	"github.com/daalfox/go-todo/todo/service"
	"github.com/daalfox/go-todo/todo/store"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	todoStore := store.NewInMemoryStore()
	todoService := service.NewService(&todoStore)

	r.Mount("/todos", &todoService)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Printf("starting server on %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not serve service: %v", err)
	}
}
