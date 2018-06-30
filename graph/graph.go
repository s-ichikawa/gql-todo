package graph

import (
	"context"
	"fmt"
	"math/rand"
	"github.com/s-ichikawa/gql-todo/model"
)

type User struct {
	ID   string
	Name string
}

type Todo struct {
	ID     string
	Text   string
	UserId string
}

type Resolver struct {
	DB model.DBModel
}

func (r *Resolver) Mutation_createUser(ctx context.Context, input NewUser) (User, error) {
	user := User{
		ID:   fmt.Sprintf("%d", rand.Int()),
		Name: input.Name,
	}

	err := r.DB.CreateUser(user.ID, user.Name)
	if err != nil {
		return User{}, fmt.Errorf("Insert error: %s", err)
	}

	return user, nil
}

func (r *Resolver) getUsers(c model.SearchUserCondition) ([]User, error) {
	rows, err := r.DB.GetUsers(c)
	if err != nil {
		return []User{}, fmt.Errorf("Select error: %s", err)
	}

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)

		users = append(users, user)
	}

	return users, nil
}

func (r *Resolver) Query_user(ctx context.Context, id string) (*User, error) {
	users, err := r.getUsers(model.SearchUserCondition{
		Ids: []string{id},
	})
	if err != nil {
		return &User{}, fmt.Errorf("Select error: %s", err)
	}

	return &users[0], nil
}

func (r *Resolver) Query_users(ctx context.Context) ([]User, error) {
	users, err := r.getUsers(model.SearchUserCondition{})
	if err != nil {
		return []User{}, fmt.Errorf("Select error: %s", err)
	}

	return users, nil
}

func (r *Resolver) Mutation_createTodo(ctx context.Context, input NewTodo) (Todo, error) {
	todo := Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserId: input.UserId,
		Text:   input.Text,
	}
	r.DB.CreateTodo(todo.ID, todo.Text, todo.UserId)
	return todo, nil
}

func (r *Resolver) getTodos(c model.SearchTodoCondition) ([]Todo, error) {
	rows, err := r.DB.GetTodos(c)
	if err != nil {
		return []Todo{}, fmt.Errorf("Select error: %s", err)
	}

	var todos []Todo
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.UserId, &todo.Text)

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *Resolver) Query_todo(ctx context.Context, id string) (*Todo, error) {
	todos, err := r.getTodos(model.SearchTodoCondition{
		Id: id,
	})
	if err != nil {
		return &Todo{}, fmt.Errorf("Select error: %s", err)
	}

	return &todos[0], nil
}

func (r *Resolver) Query_todos(ctx context.Context) ([]Todo, error) {
	todos, err := r.getTodos(model.SearchTodoCondition{})
	if err != nil {
		return []Todo{}, fmt.Errorf("Select error: %s", err)
	}
	return todos, nil
}

func (r *Resolver) Todo_user(ctx context.Context, obj *Todo) (*User, error) {
	return ctx.Value(userLoaderKey).(*UserLoader).Load(obj.UserId)
}
