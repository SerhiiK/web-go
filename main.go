package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome webgo</h1>")
}

func contactHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch email me at <a href=\"mailto:example@mail.com\">example@mail.com</a>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandle(w, r)
	default:
		// TODO: handle page not found error
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	//http.HandleFunc("/contact", contactHandle)
	fmt.Println("Starting server :3000")
	http.ListenAndServe(":3000", nil)
}
