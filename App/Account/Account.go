package account

import (
	"errors"
	"fmt"
)

func Register(database map[string]string) error {
	fmt.Print("name: ")

	name := ""
	fmt.Scan(&name)

	if _, alreadyIn := database[name]; alreadyIn {
		return errors.New("user with that name has already been registered")
	}

	fmt.Print("password: ")
	password := ""
	fmt.Scan(&password)

	database[name] = password

	return nil
}

func LogIn(database map[string]string) error {
	fmt.Print("name: ")

	name := ""
	fmt.Scan(&name)

	if _, alreadyIn := database[name]; !alreadyIn {
		return errors.New("there is no user with that name")
	}

	fmt.Print("password: ")
	password := ""
	fmt.Scan(&password)

	if password != database[name] {
		return errors.New("password is incorrect")
	}

	return nil

}

func Delete(database map[string]string) error {
	fmt.Print("name: ")

	name := ""
	fmt.Scan(&name)

	if _, alreadyIn := database[name]; !alreadyIn {
		return errors.New("there is no user with that name")
	}

	fmt.Print("password: ")

	password := ""
	fmt.Scan(&password)

	if password != database[name] {
		return errors.New("password is incorrect")
	}

	fmt.Print("Are you sure you want to delete your account?(y/n): ")

	sure := ""
	fmt.Scan(&sure)

	if sure == "y" {
		delete(database, name)
		fmt.Println("Success")

		return nil
	}

	return nil
}
