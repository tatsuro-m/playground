package main

import (
	"fmt"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/service/user"
)

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Execute SQL Statement!!")
	printHyphen()

	fmt.Println("Select")
	printHyphen()
	users, _ := user.GetAllUsers()
	for _, u := range users {
		fmt.Println(u)
	}

	u, _ := user.GetUserByName("test")
	fmt.Println(u)

	u, _ = user.GetUserByID(1)
	fmt.Println(u)

	printSharp()
	fmt.Println("insert")
	printHyphen()
	user.Insert()

	printHyphen()
}

func printHyphen() {
	fmt.Println("-------------------------------")
}

func printSharp() {
	fmt.Println("################################")
}
