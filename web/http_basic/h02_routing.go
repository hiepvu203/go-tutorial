package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetIDFromParam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := w.Write([]byte(id))
	if err != nil {
		return
	}
}

// middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request: ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Search(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	_, err := w.Write([]byte(keyword))
	if err != nil {
		return
	}
}

func main() {
	r := chi.NewRouter()

	// Gắn middleware cho toàn route
	r.Use(Logger)

	// group với middleware riêng
	r.Route("/users", func(user chi.Router) {
		user.Use(Logger)
		user.Get("/{id}", GetIDFromParam)
		user.Get("/", Search)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
