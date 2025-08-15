package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	mw "github.com/Elren44/neural-go-rest/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route ", r.URL.Path)
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	path := r.URL.Path
	userId, ok := strings.CutPrefix(path, "/teachers/")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userId = strings.TrimSuffix(userId, "/")

	fmt.Println("userId ", userId)

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method of Teachers Route"))
		fmt.Println("Hello GET Method of Teachers Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST Method of Teachers Route"))
		fmt.Println("Hello POST Method of Teachers Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method of Teachers Route"))
		fmt.Println("Hello PUT Method of Teachers Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method of Teachers Route"))
		fmt.Println("Hello DELETE Method of Teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method of Teachers Route"))
		fmt.Println("Hello PATCH Method of Teachers Route")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Method Not Allowed")
	}

}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method of Students Route"))
		fmt.Println("Hello GET Method of Students Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST Method of Students Route"))
		fmt.Println("Hello POST Method of Students Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method of Students Route"))
		fmt.Println("Hello PUT Method of Students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method of Students Route"))
		fmt.Println("Hello DELETE Method of Students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method of Students Route"))
		fmt.Println("Hello PATCH Method of Students Route")
	}

	w.Write([]byte("Hello Students Route"))
	fmt.Println("Hello Students Route ", r.URL.Path)
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method of Execs Route"))
		fmt.Println("Hello GET Method of Execs Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST Method of Execs Route"))
		fmt.Println("Hello POST Method of Execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method of Execs Route"))
		fmt.Println("Hello PUT Method of Execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method of Execs Route"))
		fmt.Println("Hello DELETE Method of Execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method of Execs Route"))
		fmt.Println("Hello PATCH Method of Execs Route")
	}

	w.Write([]byte("Hello Execs Route"))
	fmt.Println("Hello Execs Route ", r.URL.Path)
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	server := &http.Server{
		Addr:      port,
		Handler:   rl.Middleware(mw.Compression(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux))))),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port: ", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
