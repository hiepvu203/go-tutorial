package main

import (
	"fmt"
	"math"
)

const (
	MyHome  string = "Nam Dinh"
	MyPhone        = "0123456789"
)

func GetProfile() (int, string, string) {
	var age, name = 21, "Hiep"
	return age, name, MyPhone
}

func GetMyHome() string {
	return "My home " + MyHome
}

func GetMyCareer() string {
	myCareer := "Software Engineer"
	return myCareer
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(GetProfile())
	fmt.Println(GetMyHome())
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))
}
