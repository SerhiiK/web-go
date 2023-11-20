package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome webgo</h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Starting server :3000")
	http.ListenAndServe(":3000", nil)
}
