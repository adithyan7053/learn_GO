package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adithyan7053/learn_GO/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8000...")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
