package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 9000...")
	http.ListenAndServe(":9000", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
