package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.URL.String(), r.Method, time.Since(begin))
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"/api/users/{foo}",
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("foo")
			fmt.Println(id)
			fmt.Fprintln(w, "hello, world")
		},
	)

	server := &http.Server{
		Addr:                         "localhost:8080",
		Handler:                      MiddlewareLog(mux),
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Minute,
	}

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}

	fmt.Println(1)
}
