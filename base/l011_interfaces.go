package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	var s Speaker

	s = Dog{}
	fmt.Println(s.Speak())

	s = Cat{}
	fmt.Println(s.Speak())
}
