package main

import (
	"fmt"
)

func main()  {
	if cli() {
		server()
	}
	fmt.Println("Welcome to the cumzone")
}