package handlers

import (
	"github.com/gorilla/csrf"
	"net/http"
)

func CSRFProtectedHandler(w http.ResponseWriter, r *http.Request) {
	csrfToken := csrf.Token(r)
	w.Header().Set("X-CSRF-Token", csrfToken)
	w.Write([]byte("CSRF token generated"))
}
