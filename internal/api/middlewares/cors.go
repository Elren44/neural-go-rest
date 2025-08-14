package middlewares

import (
	"fmt"
	"net/http"
	"slices"
)

// AllowedOrigins is a list of origins that are allowed to make CORS requests
var allowedOrigins = []string{
	"https://localhost:3000",
	"https://test.local",
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println("Received CORS request from Origin:", origin)
		if isAllowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			http.Error(w, "CORS not allowed for this origin", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func isAllowedOrigin(origin string) bool {
	return slices.Contains(allowedOrigins, origin)
}
