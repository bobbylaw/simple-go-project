package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"simple-go-project/model"
	"simple-go-project/utils"
)

var teams map[string]model.Team
var groupRecord map[int]map[string]model.GroupRecord

func registrationPage(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}

func registerTeams(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	utils.CreateTeams(r.FormValue("registration"), &teams)
	utils.CreateGroupRecord(r.FormValue("registration"), &groupRecord)

	t, err := template.ParseFiles("./static/result.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, groupRecord)
}

func resultPage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	utils.UpdateMatchResult(r.FormValue("result"), &teams, &groupRecord)
	sortedGroup := utils.SortResult(&groupRecord)

	t, err := template.ParseFiles("./static/finalresult.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, sortedGroup)
}

func main() {
	http.HandleFunc("/", registrationPage)
	http.HandleFunc("/register", registerTeams)
	http.HandleFunc("/result", resultPage)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
