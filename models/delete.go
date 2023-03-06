package models

import "github.com/leonardogregoriocs/database"

func Delete(id int64) (int64, error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(`DELETE FROM todos WHERE id=?`, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
