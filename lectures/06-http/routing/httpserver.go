package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()

	mw := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.PathValue("id"))
			h.ServeHTTP(w, r)
		})
	}
	mux.Handle("GET /posts/{id}", mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("yes")
	})))
	r := httptest.NewRecorder()
	mux.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/posts/123", nil))
}
