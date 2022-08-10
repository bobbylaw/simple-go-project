package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"simple-go-project/config"
	"simple-go-project/model"
	"simple-go-project/utils"
)

var Teams map[string]model.Team
var GroupRecord map[int]map[string]model.GroupRecord

func RegistrationPage(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}

func RegisterTeams(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	utils.Register(r.FormValue("registration"), config.GetDB())
	groups := utils.ConvertGroups(model.GetAllGroupRecord(config.GetDB()))
	t, err := template.ParseFiles("./static/result.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, groups)
}

func ResultPage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	utils.UpdateMatchResult(r.FormValue("result"), config.GetDB())
	sortedGroup := utils.SortResult(utils.ConvertGroups(model.GetAllGroupRecord(config.GetDB())))

	t, err := template.ParseFiles("./static/finalresult.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, sortedGroup)
}
