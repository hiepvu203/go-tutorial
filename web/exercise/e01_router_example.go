package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id, "name": "User " + id})
}

/* --- middleware --- */
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("→ %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer my-secret" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()

	// Middleware toàn cục
	r.Use(middleware.Recoverer) // Tự động recover nếu panic
	r.Use(logger)

	// Health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// API Group
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", getUsers)
		r.Get("/users/{id}", getUser)

		// Admin group với xác thực
		r.Route("/admin", func(r chi.Router) {
			r.Use(authenticate)
			r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte("Welcome, Admin!"))
				if err != nil {
					return
				}
			})
		})
	})

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
