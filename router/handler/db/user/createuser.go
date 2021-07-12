package user

import (
	"log"

	db "github.com/leiendrulat/leisite/router/handler/db"
)

func CreateUser(f string, l string, e string, p string) User {
	//opening database
	data := db.Conn()
	// query
	stmt, err := data.Prepare("INSERT INTO user(fname, lname, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	userstemp := User{Fname: f, Lname: l, Email: e, Password: p}

	//add to the database
	res, err := stmt.Exec(userstemp.Fname, userstemp.Lname, userstemp.Email, userstemp.Password)
	if err != nil {
		log.Fatal(err)
	}
	//if error then print first and last id
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return userstemp
}
