package main

import "fmt"

var fruits = []string{"apple", "peach", "pear"}

func GetEachForFruit() {
	for _, value := range fruits {
		fmt.Printf("%v\n", value)
	}
}

func GetStudent() {
	studentName := []string{"Nga", "Phương", "Minh"}
	for idx, value := range studentName {
		fmt.Printf("%v\t%v\n", idx, value)
	}
}

func main() {
	GetEachForFruit()
	GetStudent()
}
