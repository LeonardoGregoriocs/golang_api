package models

import (
	"github.com/leonardogregoriocs/database"
)

func Update(id int64, todo Todo) (int64, error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec("UPDATE todos SET title = ?, description = ?, done = ? where id = ?", todo.Title, todo.Description, todo.Done, id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
