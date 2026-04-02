package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("http://localhost:3001")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		for name, values := range r.Header {
			for _, value := range values {
				log.Printf("Header: %s=%s", name, value)
			}
		}

		proxy.ServeHTTP(w, r)
	}

	log.Println("Starting proxy on :8080")
	err = http.ListenAndServe(":8080", http.HandlerFunc(handler))
	if err != nil {
		log.Fatal(err)
	}
}
