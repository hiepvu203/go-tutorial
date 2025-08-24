package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	c := make(map[string]int)
	c["Oslo"] = 1
	c["Bergen"] = 2
	c["Trondheim"] = 3
	c["Stavanger"] = 4

	fmt.Printf("c\t%v\n", c)
	c["Oslo"] = 7
	fmt.Printf("c\t%v\n", c)
	delete(c, "Oslo")
	fmt.Printf("c\t%v\n", c)

	val1, ok1 := c["Oslo"]
	val2, ok2 := c["Bergen"]
	_, ok3 := c["Trondheim"]

	fmt.Printf("val1\t%v\t%v\n", val1, ok1) // false
	fmt.Printf("val2\t%v\t%v\n", val2, ok2) // true
	fmt.Printf("val3\t%v\t\n", ok3)         // true

	for k, v := range c {
		fmt.Printf("%v\t%v\n", k, v)
	}
}
