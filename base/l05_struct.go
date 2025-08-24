package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary float32
}

type Employee struct {
	person     Person
	department string
}

func main() {
	var pers1 Person
	pers1.name = "Hege"
	pers1.age = 45
	//pers1.job = "Teacher"
	//pers1.salary = 6000
	per := Person{
		name:   "Hiep",
		age:    22,
		job:    "dev intern",
		salary: 0}

	fmt.Println(per)

	e := Employee{
		person: Person{
			name: "Hiep",
			age:  22,
			job:  "dev intern",
		},
		department: "IT",
	}

	fmt.Println(e)
}
