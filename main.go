package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)

	http.HandleFunc("/add", handleAdd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start:", err)
	}
}
