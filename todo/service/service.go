package service

import (
	"github.com/daalfox/go-todo/todo/store"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Service struct {
	router chi.Router
	repo   store.Store
}

func NewService(repo store.Store) Service {
	s := Service{
		repo:   repo,
		router: chi.NewRouter(),
	}
	s.addRoutes()
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Service) addRoutes() {
	r := s.router
	r.Get("/", s.listTodos)
	r.Post("/", s.addTodo)
	r.Get("/{id}", s.getTodo)
	r.Delete("/{id}", s.deleteTodo)
	r.Put("/{id}", s.updateTodo)
}
