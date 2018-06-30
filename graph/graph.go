package graph

import (
	"context"
	"fmt"
	"math/rand"
	"github.com/pkg/errors"
)

type User struct {
	ID   string
	Name string
}

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

type MyApp struct {
	users []User
	todos []Todo
}

func (a *MyApp) Query_todo(ctx context.Context, id string) (Todo, error)  {
	for _, todo := range a.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return Todo{}, errors.New("Not Found")
}

func (a *MyApp) Query_todos(ctx context.Context) ([]Todo, error) {
	return a.todos, nil
}

func (a *MyApp) Mutation_createTodo(ctx context.Context, input NewTodo) (Todo, error) {
	todo := Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: fmt.Sprintf("U%d", input.UserId),
		Text:   input.Text,
	}
	a.todos = append(a.todos, todo)
	return todo, nil
}

func (a *MyApp)	Mutation_createUser(ctx context.Context, input NewUser) (User, error)  {
	user := User{
		ID: fmt.Sprintf("%d", rand.Int()),
		Name:input.Name,
	}
	a.users = append(a.users, user)
	return user, nil
}

func (a *MyApp) Todo_user(ctx context.Context, it *Todo) (User, error) {
	return User{ID: it.UserID, Name: "user " + it.UserID}, nil
}
