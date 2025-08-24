package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	// --- POINTERS TO STRUCTS ---
	v := Vertex{1, 2}
	ptr := &v
	ptr.X = 1e9
	fmt.Println(v) // {1000000000 2}

	//fmt.Println(v1, point, v2, v3)
}
