package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// If the client does not support gzip, skip compression
			next.ServeHTTP(w, r)
			return
		}

		// Set the Content-Encoding header to gzip
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		w = &gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gz,
		}

		next.ServeHTTP(w, r)
		fmt.Println("Response compressed with gzip for", r.URL.Path)
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (grw *gzipResponseWriter) Write(b []byte) (int, error) {
	if grw.Writer == nil {
		return grw.ResponseWriter.Write(b)
	}
	return grw.Writer.Write(b)
}
