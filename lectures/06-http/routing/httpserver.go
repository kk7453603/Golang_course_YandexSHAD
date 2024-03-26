package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.PathValue("id"))
	})
	r := httptest.NewRecorder()
	mux.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/posts/123", nil))
}
