package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leonardogregoriocs/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)

	conn, err := sql.Open("mysql", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
