package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var r Rectangle
	fmt.Print("Enter width and height: ")
	_, err := fmt.Scan(&r.width, &r.height)
	if err != nil {
		return
	}
	fmt.Print(r.Area())
	fmt.Println(r.Perimeter())
}
