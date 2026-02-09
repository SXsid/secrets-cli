package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world from my app"))
	})
	log.Printf("Sever is up at port :%d", 8080)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
