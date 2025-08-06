package store

import (
	"errors"
	"slices"

	"github.com/daalfox/go-todo/todo"
)

var (
	ErrNotFound = errors.New("item not found")
)

type InMemoryStore struct {
	todos   []*todo.Todo
	last_id int
}

func NewInMemoryStore() InMemoryStore {
	return InMemoryStore{todos: make([]*todo.Todo, 0)}
}

func (s *InMemoryStore) List() ([]*todo.Todo, error) {
	return s.todos, nil
}
func (s *InMemoryStore) Get(id int) (*todo.Todo, error) {
	for _, item := range s.todos {
		if item.Id == id {
			return item, nil
		}
	}

	return nil, ErrNotFound
}
func (s *InMemoryStore) Add(title string) (*todo.Todo, error) {
	next := s.last_id + 1
	newTodo := &todo.Todo{
		Id:    next,
		Title: title,
	}
	s.todos = append(s.todos, newTodo)

	s.last_id = next
	return newTodo, nil
}
func (s *InMemoryStore) Delete(id int) error {
	for i, item := range s.todos {
		if item.Id == id {
			s.todos = slices.Delete(s.todos, i, i+1)
			return nil
		}
	}

	return ErrNotFound
}
func (s *InMemoryStore) Update(id int, updated todo.Todo) error {
	var target *todo.Todo
	for _, item := range s.todos {
		if item.Id == id {
			target = item
		}
	}
	if target == nil {
		return ErrNotFound
	}
	target.Title = updated.Title
	target.Done = updated.Done

	return nil
}
