package main

import (
	"log"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "Hello, %s.", r.URL.Path[1:]) 
	log.Println("GET", r.URL.Path)
}

func main() {
	http.ListenAndServe(":3434", http.HandlerFunc(handler))
}
