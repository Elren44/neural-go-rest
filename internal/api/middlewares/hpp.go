package middlewares

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

type HPPOptions struct {
	CheckQuery                  bool
	CheckBody                   bool
	CheckBodyOnlyForContentType string
	WhiteList                   []string
}

func Hpp(option HPPOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if option.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, option.CheckBodyOnlyForContentType) {
				// filter the body params
				filterBodyParams(r, option.WhiteList)
			}
			if option.CheckQuery && r.URL.Query() != nil {
				// filter the query params
				filterQueryParams(r, option.WhiteList)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func filterBodyParams(r *http.Request, whiteList []string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0])
			// r.Form.Set(k, v[len(v) - 1])
		}
		if !isWhiteListed(k, whiteList) {
			delete(r.Form, k)
		}
	}
}

func filterQueryParams(r *http.Request, whiteList []string) {
	query := r.URL.Query()

	for k, v := range query {
		if len(v) > 1 {
			query.Set(k, v[0])
			// query.Set(k, v[len(v) - 1])
		}
		if !isWhiteListed(k, whiteList) {
			query.Del(k)
		}
	}
	r.URL.RawQuery = query.Encode()
}

func isWhiteListed(param string, whiteList []string) bool {
	return slices.Contains(whiteList, param)
}
