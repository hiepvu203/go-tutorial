package main

import (
	"fmt"
)

func main() {
	// --- LESSON 1: VARIABLES ---
	fmt.Println(GetProfile())
	fmt.Println(GetMyHome())

	// --- LESSON 2: ARRAYS
	fmt.Println(GetMyTypes())
	fmt.Println(GetMyScores())
	fmt.Println(CreateSliceFromArray())
	AppendSlice()

	// --- LESSON 3: LOOPS
	GetEachForFruit()
	GetStudent()

	// --- LESSON 4: FUNCTION
	sum := sum(1, 2)
	fmt.Println(sum)

	multiple := multiple(3, 4)
	fmt.Println(multiple)

	age, name := GetInfo(22, "Hiep")
	fmt.Println(age, name)
}
