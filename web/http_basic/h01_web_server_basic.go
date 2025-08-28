package main

import (
	"fmt"
	"io"
	"net/http"
)

type User struct {
}

func GetName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "my-value")

	_, err := w.Write([]byte("Hello " + name))
	if err != nil {
		return
	}
}

func GetMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("Method GET"))
	case "POST":
		w.Write([]byte("Method POST"))
	case "PUT":
		w.Write([]byte("Method PUT"))
	case "DELETE":
		w.Write([]byte("Method DELETE"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

func LoadDataForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	age := r.FormValue("age")
	if age == "" {
		age = "0"
	}
	_, err = w.Write([]byte(age))
}

func ReadBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Không thể đọc body", http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body) // luôn đóng body sau khi dùng

	_, err = w.Write(body)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello World")
		if err != nil {
			return
		}
	})

	http.HandleFunc("/profile", GetName)
	http.HandleFunc("/method", GetMethod)
	http.HandleFunc("/form", LoadDataForm)
	http.HandleFunc("/read", ReadBody)

	fmt.Println("Server running at: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
