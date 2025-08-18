package main

import "fmt"

func GetMyTypes() []string {
	myTypes := []string{"confidence", "kindness"}
	return myTypes
}

var MyScores = []float32{9.8, 7.9}

func GetMyScores() []float32 {
	MyScores[1] = 8.9
	return MyScores
}

var arr1 = [6]int{1, 2, 3, 4, 5, 6}

func CreateSliceFromArray() []int {
	newSlice := arr1[1:4]
	result := []int{
		len(newSlice),
		cap(newSlice),
	}
	return result
}

func AppendSlice() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(myslice1), cap(myslice1))

	myslice1 = append(myslice1, 10, 21)
	fmt.Println(len(myslice1), cap(myslice1))
}
