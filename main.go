package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmid"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "text/html")
	
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello World Dynamic oke</h1>")	
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<a href=\"mailto:arief.hasnul@gmail.com\">Contact us</a>")	
	} else {
		w.WriteHader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 Page not found</h1>")
	}
}

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	//http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}