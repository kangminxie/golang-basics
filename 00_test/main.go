package main

import "fmt"

const (
	LANGUAGE = "Golang"
	YEAR     = 2023
)

func main() {
	firstName, lastName := "MK", "Developer"
	fmt.Printf("Hello %s %s\n", firstName, lastName)
	fmt.Printf("We are learning %s programming at %d", LANGUAGE, YEAR)
}
