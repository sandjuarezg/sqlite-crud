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
			var err = user.ReadData()
			if err != nil {
				fmt.Println("Problem: Table blank")
			}
			function.CleanConsole()
		case 3:
			var n = user.UpdateData()
			if n == -1 {
				fmt.Println("Problem: Table blank")
			}
			if n == 0 {
				fmt.Println("Problem: Not found id")
			}
			function.CleanConsole()
		case 4:
			var n = user.DeleteData()
			if n == -1 {
				fmt.Println("Problem:  Table blank")
			}
			if n == 0 {
				fmt.Println("Problem: Not found id")
			}
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
