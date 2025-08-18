package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1    = Vertex{1, 2}
	v2    = Vertex{X: 1}
	v3    = Vertex{}
	point = &Vertex{1, 2}
)

func main() {
	// --- BASE ---
	i, j := 42, 2701
	p := &i         // trỏ tới i
	fmt.Println(*p) // đọc i thông qua con trỏ
	*p = 21         // set giá trị của i thông qua con trỏ
	fmt.Println(i)  // giá trị mới của i

	p = &j         // trỏ tới j
	*p = *p / 37   // chia j thông qua con trỏ
	fmt.Println(j) // giá trị mới của j

	// --- POINTERS TO STRUCTS ---
	v := Vertex{1, 2}
	ptr := &v
	ptr.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, point, v2, v3)
}
