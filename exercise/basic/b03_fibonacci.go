package main

import "fmt"

func fibonacci(n int) {
	if n <= 1 {
		fmt.Println(n)
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
		fmt.Println(a, b)
	}
	fmt.Println(b)
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fibonacci(n)
}
