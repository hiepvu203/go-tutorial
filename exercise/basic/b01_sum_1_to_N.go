package main

import "fmt"

func sum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fmt.Println("Sum:", sum(n))
}
