package models

import (
	"context"
	"log"

	"github.com/leonardogregoriocs/database"
)

func Insert(todo Todo) (int64, error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := "INSERT INTO `todos` (`title`, `description`, `done`) VALUES (?, ?, ?)"

	result, err := conn.ExecContext(context.Background(), sql, todo.Title, todo.Description, todo.Done)

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Erro ao inserir registro: %s", err)
	}

	return id, err
}
