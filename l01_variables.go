package main

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
