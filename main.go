package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	conf := readConfig()

	listenAt := fmt.Sprintf(":%d", conf.port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responseFormat := r.URL.Query().Get("format")
		if responseFormat == "json" {
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			json.NewEncoder(w).Encode(struct {
				Foo string `json:"foo"`
			}{
				Foo: "Hello, World!",
			})
			return
		}
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "bar")
	})

	log.Printf("Open the following URL in the browser: http://localhost:%d\n", conf.port)
	log.Fatal(http.ListenAndServe(listenAt, nil))
}
