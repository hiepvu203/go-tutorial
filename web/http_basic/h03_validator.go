package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validateAndRespond(v *validator.Validate, data interface{}, w http.ResponseWriter) bool {
	if err := v.Struct(data); err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			field := lowerFirst(e.Field())
			errors[field] = getErrorMessage(e)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]map[string]string{"errors": errors})
		return false
	}
	return true
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "trường này là bắt buộc"
	case "email":
		return "phải là email hợp lệ"
	case "min":
		return "phải có ít nhất " + e.Param() + " ký tự"
	case "max":
		return "không được vượt quá " + e.Param() + " ký tự"
	default:
		return "không hợp lệ"
	}
}
