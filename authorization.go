package main

import (
	"fmt"
	"net/http"
)

var apiToken string

func authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if apiToken == "" {
			return
		}

		key := firstOrDefault(r.URL.Query()["apiToken"])
		fmt.Println(key)
		if key == "" {
			key = r.Header.Get("apiToken")
			fmt.Println(key)
		}
		if key == apiToken {
			next.ServeHTTP(w, r)
		} else {
			w.Write([]byte("Access denied"))
			w.WriteHeader(http.StatusForbidden)
		}
	})
}

func firstOrDefault(slice []string) string {
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}
