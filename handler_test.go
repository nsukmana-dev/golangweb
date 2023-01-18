package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Nandar")
	}

	server := http.Server{
		Addr:    "localhost:8790",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Index")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})
	mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Image")
	})
	mux.HandleFunc("/image/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8790",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
