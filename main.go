package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("Hello World"))
		data, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("oops"))
			return
		}
		w.Write([]byte("Hello " + string(data)))
	})
	http.ListenAndServe(":8080", nil)
}
