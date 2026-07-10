package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"io"
)

func main() {
	fs := http.FileServer(http.Dir("./PayloadInterface"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	res := w.Result()
	body, _ := io.ReadAll(res.Body)
	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body length: %d\n", len(body))

    req2 := httptest.NewRequest("GET", "/style/style.css", nil)
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	res2 := w2.Result()
	body2, _ := io.ReadAll(res2.Body)
	fmt.Printf("Status: %d\n", res2.StatusCode)
	fmt.Printf("Body length: %d\n", len(body2))
}
