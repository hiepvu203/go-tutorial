package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Person struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,min=2,max=50"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var persons = []Person{
	{ID: 1, Name: "Alice", Email: "alice@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Name: "Bob", Email: "bob@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// middleware
func withJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func GetID(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID không hợp lệ", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(persons)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	id, ok := GetID(w, r)
	if !ok {
		return
	}

	for _, person := range persons {
		if person.ID == id {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	http.Error(w, "Person not found", http.StatusNotFound)
}

var validate = validator.New()

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var NewPerson Person
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Không thể đọc body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &NewPerson); err != nil {
		http.Error(w, "Json không hợp lệ", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(NewPerson); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}

		json.NewEncoder(w).Encode(errors)
		return
	}
	NewPerson.ID = len(persons) + 1
	NewPerson.CreatedAt = time.Now()
	NewPerson.UpdatedAt = time.Now()
	persons = append(persons, NewPerson)
	json.NewEncoder(w).Encode(persons)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, ok := GetID(w, r)
	if !ok {
		return
	}

	var UpdatePerson Person
	if err := json.NewDecoder(r.Body).Decode(&UpdatePerson); err != nil {
		http.Error(w, "JSON khong hop le", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(UpdatePerson); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}

		json.NewEncoder(w).Encode(errors)
		return
	}

	for i, user := range persons {
		if user.ID == id {
			UpdatePerson.ID = id
			UpdatePerson.CreatedAt = user.CreatedAt
			UpdatePerson.UpdatedAt = time.Now()
			persons[i] = UpdatePerson

			json.NewEncoder(w).Encode(UpdatePerson)
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	id, ok := GetID(w, r)
	if !ok {
		return
	}

	for i, person := range persons {
		if person.ID == id {
			persons = append(persons[:i], persons[i+1:]...)

		}
	}
}

func main() {
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(withJSON)
		r.Get("/persons", GetPersons)
		r.Get("/persons/{id}", GetPerson)
		r.Post("/persons", CreatePerson)
	})

}
