package main

import (
	"log"
	"net/http"

	handler "piscine/handlers"
)

func main() {
	file := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", file))

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/ascii", handler.AsciiArtHandler)
	http.HandleFunc("/404",handler.NotFoundHandler)
	
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.NotFound(w, r)
	// })
	log.Println("http://localhost:8082")
	http.ListenAndServe(":8082", nil)
}
