package handlers

import (
	"fmt"
	"net/http"
)

// Protected handles GET /protected — requires valid JWT
func Protected(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a protected endpoint! You are authenticated.")
}
