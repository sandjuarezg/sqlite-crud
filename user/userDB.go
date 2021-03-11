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
	active   bool
}

func InsertData() {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, username TEXT, pass TEXT, active BIT)")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = db.Prepare("INSERT INTO users (name, username, pass, active) VALUES (?, ?, ?, ?)")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec("Marco Diaz", "marco123", "passMarco", 1)
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
		rows.Scan(&user.id, &user.name, &user.username, &user.pass, &user.active)
		fmt.Println(user.id, user.name, user.username, user.pass, user.active)
	}
}

func UpdateData() {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var statement, err = db.Prepare("UPDATE users SET username = ? WHERE id = ?")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	var res, _ = statement.Exec("aaaaaaaaaaaaaaa", 2)
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
	var res, _ = statement.Exec(2)
	res.RowsAffected()
}
