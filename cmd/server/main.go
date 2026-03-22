package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting zs-int-godata ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
