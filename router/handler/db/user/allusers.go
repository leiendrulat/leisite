package user

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/leiendrulat/leisite/router/handler/db"
)

type User struct {
	ID       string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Fname    string `param:"fname" query:"fname" form:"fname" json:"fname" xml:"fname"`
	Lname    string `param:"lname" query:"lname" form:"lname" json:"lname" xml:"lname"`
	Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
	Password string `param:"password" query:"password" form:"password" json:"password" xml:"password"`
}

//CreateUser creates a Post
func GetAllDataUser() []User {
	//opening database
	data := db.Conn()

	var (
		id       string
		fname    string
		lname    string
		email    string
		password string
		user     []User
	)
	i := 0
	//get from database
	rows, err := data.Query("select id, fname, lname, email, password from user")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &fname, &lname, &email, &password)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}
		u := User{ID: id, Fname: fname, Lname: lname, Email: email, Password: password}
		user = append(user, u)

	}
	defer rows.Close()
	defer data.Close()
	return user
}

//CreateUser creates a Post
func GetUser(ids int64) (string, string, string) {
	//opening database
	data := db.Conn()

	var (
		id    string
		fname string
		lname string
		email string
	)
	stmt, err := data.Prepare("SELECT id, fname, lname, email FROM user WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(ids).Scan(&id, &fname, &lname, &email)
	fmt.Println(err)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		// close db when not in use
		defer data.Close()
		return fname, lname, id

	case nil:
		fmt.Println("was nil", fname, lname)

		// close db when not in use
		defer data.Close()
		return fname, lname, id

	default:

		fmt.Println("default!!!!!!!!!!!!")

		return fname, lname, id
	}
}
