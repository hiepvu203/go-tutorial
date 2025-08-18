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

	// range
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i  --  toán tử dịch bit trái
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
