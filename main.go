package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-go-project/routes"
)

func main() {
	routes.RegisterApplicationRoutes()
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
