package main

import (
	"fmt"
)

func cli() bool {
	var choice string
	fmt.Println("Welcome to goDZilla, a simple file server with minimal configuration")
	fmt.Println("Do you want to start the server? ")
	fmt.Scan(&choice)
	return choice=="yes"
}