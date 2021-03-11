package main

import (
	"fmt"

	"github.com/sandjuarezg/sqlite-crud/function"

	"github.com/sandjuarezg/sqlite-crud/user"
)

func main() {
	var opc int
	var exit bool

	for !exit {
		fmt.Println("-- What would you like to do? --")
		fmt.Println("1. Insert data")
		fmt.Println("2. Read data")
		fmt.Println("3. Update data")
		fmt.Println("4. Delate data")
		fmt.Println("5. Exit")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			user.InsertData()
			function.CleanConsole()
		case 2:
			user.ReadData()
			function.CleanConsole()
		case 3:
			user.UpdateData()
			function.CleanConsole()
		case 4:
			user.DeleteData()
			function.CleanConsole()
		case 5:
			fmt.Println("E X I T . . .")
			exit = true
		default:
			fmt.Println("Option not valid")
			function.CleanConsole()
		}
	}
}
