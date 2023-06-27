package main

import (
	"net/http"
)

func hellohangmanHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", hellohangmanHandler)
	http.ListenAndServe(":2555", nil)
}
