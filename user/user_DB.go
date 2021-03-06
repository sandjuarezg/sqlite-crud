package user

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id       int
	name     string
	username string
	pass     string
}

func SqlMigration() {
	//Check migraton.sql
	var _, err = os.Stat("./migration.sql")
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	//Get content
	file, _ := os.Open("./migration.sql")
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Check userdata.db
	_, err = os.Stat("./user/userdata.db")
	if os.IsNotExist(err) {
		var file, err = os.Create("./user/userdata.db")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

	}

	//Check table
	db, err := sql.Open("sqlite3", "./user/userdata.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Query("SELECT * FROM users")
	if err != nil {
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func InsertData() {
	var db, err = sql.Open("sqlite3", "./user/userdata.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (name, username, pass) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

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
	var db, err = sql.Open("sqlite3", "./user/userdata.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, username, pass FROM users")
	if err != nil {
		return err
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
	var db, err = sql.Open("sqlite3", "./user/userdata.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Query("SELECT * FROM users")
	if err != nil {
		return -1
	}

	statement, err := db.Prepare("UPDATE users SET pass = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

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
	var db, err = sql.Open("sqlite3", "./user/userdata.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Query("SELECT * FROM users")
	if err != nil {
		return -1
	}

	statement, err := db.Prepare("DELETE from users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var user = user{}
	fmt.Println("Enter id")
	fmt.Scan(&user.id)
	var res, _ = statement.Exec(user.id)
	var n, _ = res.RowsAffected()

	return n
}
