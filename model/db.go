package model

import (
	"database/sql"
	"fmt"
	"strings"
)

type DBModel struct {
	DB *sql.DB
}

type SearchUserCondition struct {
	Ids []string
}

type SearchTodoCondition struct {
	Id string
}

func (m *DBModel) CreateUser(id string, name string) (error) {
	_, err := m.DB.Exec("INSERT INTO users (id, name) VALUES (?, ?)", id, name)
	if err != nil {
		return fmt.Errorf("Insert users error: %s", err)
	}
	return nil
}

func (r *DBModel) GetUsers(c SearchUserCondition) (*sql.Rows, error) {
	query := "SELECT id, name FROM users "

	args := make([]interface{}, len(c.Ids))
	if len(c.Ids) > 0 {
		placeholders := make([]string, len(c.Ids))
		for i := 0; i < len(c.Ids); i++ {
			placeholders[i] = "?"
			args[i] = c.Ids[i]
		}

		query += " WHERE id IN (" + strings.Join(placeholders, ",") + ")"
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Select users error: %s", err)
	}

	return rows, nil
}

func (m *DBModel) CreateTodo(id string, text string, userId string) (error) {
	_, err := m.DB.Exec("INSERT INTO todos (id, user_id, text) VALUES (?, ?, ?)", id, userId, text)
	if err != nil {
		return fmt.Errorf("Insert todos error: %s", err)
	}
	return nil
}

func (r *DBModel) GetTodos(c SearchTodoCondition) (*sql.Rows, error) {
	query := "SELECT id, user_id, text FROM todos "

	var args []interface{}
	if len(c.Id) > 0 {
		query += " WHERE id = ?"
		args = append(args, c.Id)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Select todos error: %s", err)
	}

	return rows, nil
}
