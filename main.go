package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello ")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this program have beeen dockerized!"))
	})

	http.ListenAndServe(":3000", nil);
}
