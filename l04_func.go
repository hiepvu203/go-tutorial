package main

func sum(x, y int) int {
	return x + y
}

func multiple(x, y int) (result int) {
	result = x * y
	return
}
func GetInfo(age int, name string) (result int, text string) {
	result = age
	text = "His name is " + name
	return
}
