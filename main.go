package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,Geekbrains")
}

func main() {
	http.HandleFunc("/", echoFunction)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	log.Fatal(http.ListenAndServe(":8085", nil))

}
