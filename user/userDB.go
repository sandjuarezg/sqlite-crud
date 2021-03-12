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

func ReadData() error {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var rows, errSelect = db.Query("SELECT * FROM users")
	if errSelect != nil {
		return errSelect
	}
	defer rows.Close()

	var user = user{}
	fmt.Printf("|%-6s|%-15s|%-15s|%-15s|\n", "id", "Name", "User name", "Password")
	fmt.Println("________________________________________________________")
	for rows.Next() {
		rows.Scan(&user.id, &user.name, &user.username, &user.pass)
		fmt.Printf("|%-6d|%-15s|%-15s|%-15s|\n", user.id, user.name, user.username, user.pass)
	}

	return nil
}

func UpdateData() int64 {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var _, errSelect = db.Query("SELECT * FROM users")
	if errSelect != nil {
		return -1
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
	var n, _ = res.RowsAffected()

	return n
}

func DeleteData() int64 {
	var db, errOpen = sql.Open("sqlite3", "./user/userdata.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	var _, errSelect = db.Query("SELECT * FROM users")
	if errSelect != nil {
		return -1
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
	var n, _ = res.RowsAffected()

	return n
}
