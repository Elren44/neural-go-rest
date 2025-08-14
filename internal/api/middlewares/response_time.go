package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrippedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		duration := time.Since(start)
		wrippedWriter.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrippedWriter, r)
		duration = time.Since(start)
		fmt.Printf("Method: %s, URL: %s, Status: %d, Duration: %v \n", r.Method, r.URL.Path, wrippedWriter.status, duration.String())
		fmt.Println("Send Response from ResponseTimeMiddleware")
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
