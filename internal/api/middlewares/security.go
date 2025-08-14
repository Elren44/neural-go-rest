package middlewares

import (
	"net/http"
)

func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("X-Frame-Options", "DENY ")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains;preload")
		w.Header().Set("Content-Security-Policy", "self")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("X-Powered-By", "Go Web Server")
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		w.Header().Set("Permissions-Policy", "geolocation=(self), microphone=()")

		next.ServeHTTP(w, r)
	})
}

//BasicMiddlewareSkeleton
// func BasicMiddlewareSkeleton(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Implement basic middleware logic here
// 		next.ServeHTTP(w, r)
// 	})
// }
