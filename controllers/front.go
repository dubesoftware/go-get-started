package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers maps URL patterns to handlers.
func RegisterControllers() {
	uc := NewUserController()

	http.Handle(`/users`, *uc)
	http.Handle(`/users/`, *uc)
}

// EncodeResponseAsJSON converts a user object to JSON
func EncodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
