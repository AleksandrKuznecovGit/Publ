package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,Geekbrains")
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	log.Fatal(http.ListenAndServe(":8085", nil))

}
