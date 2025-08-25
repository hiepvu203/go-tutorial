package main

import "fmt"

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func doSomething() error {
	return &MyError{Code: 404, Message: "Not Found"}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err) // "Code 404: Not Found"
	}
}
