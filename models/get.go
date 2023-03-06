package models

import "github.com/leonardogregoriocs/database"

func Get(id int64) (todo Todo, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM `todos` WHERE id=?", id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
