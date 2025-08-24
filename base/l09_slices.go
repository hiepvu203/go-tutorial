package main

import (
	"fmt"
)

func addElement() {
	s := make([]int, 0, 1)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	}
}

func main() {
	// --- slices of slices
	//board := [][]string{
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//}
	//
	//board[0][0] = "X"
	//board[2][2] = "O"
	//board[1][2] = "X"
	//board[1][0] = "O"
	//board[0][2] = "X"
	//
	//for i := 0; i < len(board); i++ {
	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
	//}
	//
	//s := make([]int, 3)
	//s[0] = 1
	//s[1] = 2
	//s[2] = 3
	//fmt.Println(s)
	//
	//clothes := make([]string, 5, 10)
	//clothes[0] = "X"
	//clothes[1] = "O"
	//clothes[2] = "X"
	//fmt.Println(clothes)

	addElement()
}
