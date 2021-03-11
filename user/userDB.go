package user

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id       int
	name     string
	username string
	pass     string
}

func InsertData() {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, username TEXT, pass TEXT)")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = db.Prepare("INSERT INTO users (name, username, pass) VALUES (?, ?, ?)")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	var user = user{}
	fmt.Println("Enter a name")
	fmt.Scan(&user.name)
	fmt.Println("Enter a username")
	fmt.Scan(&user.username)
	fmt.Println("Enter a password")
	fmt.Scan(&user.pass)

	statement.Exec(user.name, user.username, user.pass)
}

func ReadData() {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var rows, errSelect = db.Query("SELECT * FROM users")
	defer rows.Close()
	if errSelect != nil {
		log.Fatal(errSelect)
	}

	var user = user{}
	for rows.Next() {
		rows.Scan(&user.id, &user.name, &user.username, &user.pass)
		fmt.Printf("|%-6d|%-15s|%-15s|%-15s|\n", user.id, user.name, user.username, user.pass)
	}
}

func UpdateData() {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var statement, err = db.Prepare("UPDATE users SET pass = ? WHERE id = ?")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}

	var user = user{}
	fmt.Println("Enter id")
	fmt.Scan(&user.id)
	fmt.Println("Enter password to update")
	fmt.Scan(&user.pass)

	var res, _ = statement.Exec(user.pass, user.id)
	res.RowsAffected()
}

func DeleteData() {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var statement, err = db.Prepare("DELETE from users WHERE id = ?")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}

	var user = user{}
	fmt.Println("Enter id")
	fmt.Scan(&user.id)
	var res, _ = statement.Exec(user.id)
	res.RowsAffected()
}
