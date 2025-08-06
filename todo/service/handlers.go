package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/daalfox/go-todo/pkg/utils"
	"github.com/daalfox/go-todo/todo"
	"github.com/daalfox/go-todo/todo/store"
	"github.com/go-chi/chi/v5"
)

func (s *Service) listTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := s.repo.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	utils.Encode(w, http.StatusOK, todos)
}

func (s *Service) addTodo(w http.ResponseWriter, r *http.Request) {
	newTodo, err := utils.Decode[*todo.Todo](r)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode todo: %v", err), http.StatusBadRequest)
		return
	}
	problems := newTodo.Validate()
	if len(problems) > 0 {
		utils.Encode(w, http.StatusBadRequest, problems)
		return
	}

	newTodo, err = s.repo.Add(newTodo.Title)
	utils.Encode(w, http.StatusCreated, newTodo)
}

func (s *Service) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "could not parse id", http.StatusBadRequest)
		return
	}

	todo, err := s.repo.Get(id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to fetch todo: %v", err)
		return
	}

	utils.Encode(w, http.StatusOK, todo)
}

func (s *Service) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "could not parse id", http.StatusBadRequest)
		return
	}

	err = s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to fetch todo: %v", err)
		return
	}
}

func (s *Service) updateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "could not parse id", http.StatusBadRequest)
		return
	}
	todo, err := utils.Decode[todo.Todo](r)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode todo: %v", err), http.StatusBadRequest)
		return
	}

	err = s.repo.Update(id, todo)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to fetch todo: %v", err)
		return
	}

	utils.Encode(w, http.StatusOK, todo)
}
