package main

import (
	"fmt"

	"unit-test/services"
)

func main() {

	add := services.Add(2, 2)

	fmt.Println("It works", add)
}
