package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	var s string
	fmt.Print("Enter a string: ")
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}

	fmt.Println("reverse string:", reverseString(s))
}
