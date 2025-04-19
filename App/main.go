package main

import (
	"fmt"

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
			fmt.Print("name: ")

			name := ""
			fmt.Scan(&name)

			if _, alreadyIn := database[name]; alreadyIn {
				fmt.Println("User with that name has already been registered")
				break
			}

			fmt.Print("password: ")
			password := ""
			fmt.Scan(&password)

			database[name] = password

			fmt.Println("Success")

		case 2:
			fmt.Print("name: ")

			name := ""
			fmt.Scan(&name)

			if _, alreadyIn := database[name]; !alreadyIn {
				fmt.Println("There is no user with that name")
				break
			}

			fmt.Print("password: ")
			password := ""
			fmt.Scan(&password)

			if password != database[name] {
				fmt.Println("password is incorrect")
				break
			}

			fmt.Println("Success")

		case 3:
			err := storage.UpdateDatabase(fileName, database)
			if err != nil {
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
	fmt.Println("3. Exit")

	fmt.Print("choice: ")
}
