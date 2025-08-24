package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

	//file, err := os.Open("data.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//
	//	}
	//}(file)
	//
	//fmt.Println("Đã đọc file")

	//fmt.Println("counting")
	//
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}

	//fmt.Println("done")
}
