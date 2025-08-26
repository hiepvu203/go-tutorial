package main

import (
	"fmt"
	"strings"
)

func countCharacters(s string) {
	if len(s) == 0 {
		return
	}
	s = strings.ToLower(strings.TrimSpace(s))
	tmp := make(map[rune]int)
	for _, char := range s {
		tmp[char]++
	}
	var keys []rune
	for k := range tmp {
		keys = append(keys, k)
	}

	//sort.Slice(keys, func(i, j int) bool {
	//	return keys[i] < keys[j]
	//})

	for _, char := range keys {
		fmt.Printf("%c: %d ", char, tmp[char])
	}
	fmt.Println()
}

func main() {
	var s string
	fmt.Print("Enter a string: ")
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}

	countCharacters(s)
}
