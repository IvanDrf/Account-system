package main

import (
	"fmt"

	account "github.com/IvanDrf/Account-system/App/Account"
	storage "github.com/IvanDrf/Account-system/App/Storage"
)

const fileName = "../Data.json"

func main() {
	database, err := storage.ReadDatabase(fileName)

	if err != nil {
		fmt.Println(err)

		return
	}

	for {
		WriteMenu()

		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if err := account.Register(database); err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println("Success")

		case 2:
			if err := account.LogIn(database); err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println("Success")

		case 3:
			err := account.Delete(database)
			if err != nil {
				fmt.Println(err)
			}

		case 4:
			if err := storage.UpdateDatabase(fileName, database); err != nil {
				fmt.Println(err)
			}

			return

		default:
			fmt.Println("There are only three options")
		}

		fmt.Println()
	}
}

func WriteMenu() {
	fmt.Println("1. Register")
	fmt.Println("2. Log in")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")

	fmt.Print("choice: ")
}
