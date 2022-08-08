package main

import (
	"fmt"
	"log"
	"net/http"
)

func registrationPage(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}

func main() {
	http.HandleFunc("/", registrationPage)
	//http.HandleFunc("/result", resultPage)
	//http.HandleFunc("/register", registerTeams)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
