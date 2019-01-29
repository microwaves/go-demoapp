package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var port = 8080

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Heey! This is just a demo page. :-)")
}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Now listening on port", fmt.Sprintf("%d", port))
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", port), nil))
}
