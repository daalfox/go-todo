package store

import "github.com/daalfox/go-todo/todo"

// Store defines the interface for accessing and modifying todos.
//
// Methods return pointers to `Todo`s to allow distinguishing between
// non-existent and zero-valued todos.
type Store interface {
	List() ([]*todo.Todo, error)
	Get(int) (*todo.Todo, error)
	Add(string) (*todo.Todo, error)
	Delete(int) error
	Update(int, todo.Todo) error
}
