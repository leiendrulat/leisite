package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//start connection to db
func Conn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/user")
	if err != nil {
		fmt.Println(err)
	}

	//check if it pings
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return db
}
