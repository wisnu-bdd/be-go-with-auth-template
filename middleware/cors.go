package middleware

import (
	"net/http"

	// "github.com/wisnu-bdd/be-go-with-auth-template/config"
)

// CORS applies CORS headers to the response and handles preflight requests.
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// // Allow only configured origins
		// allowed := false
		// for _, allowedOrigin := range config.AllowedOrigins {
		// 	if origin == allowedOrigin {
		// 		allowed = true
		// 		break
		// 	}
		// }
		// if allowed {
		// 	w.Header().Set("Access-Control-Allow-Origin", origin)
		// }

		// // To allow any origin, comment out the above origin check and use this instead:
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		
		// Or if you need to support credentials with any origin:
		w.Header().Set("Access-Control-Allow-Origin", origin)
		// Note: Using credentials with wildcard (*) origin is not allowed by browsers

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
