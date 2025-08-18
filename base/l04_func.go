package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func multiple(x, y int) (result int) {
	result = x * y
	return
}
func GetInfo(age int, name string) (result int, text string) {
	result = age
	text = "His name is " + name
	return
}

func main() {
	sum := sum(1, 2)
	fmt.Println(sum)

	multiple := multiple(3, 4)
	fmt.Println(multiple)

	age, name := GetInfo(22, "Hiep")
	fmt.Println(age, name)
}
