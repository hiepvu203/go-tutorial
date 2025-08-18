package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary float32
}

func main() {
	var pers1 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	fmt.Println(pers1)
}
