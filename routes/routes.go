package routes

import (
	"net/http"
	"simple-go-project/controller"
)

var RegisterApplicationRoutes = func() {
	http.HandleFunc("/", controller.RegistrationPage)
	http.HandleFunc("/register", controller.RegisterTeams)
	http.HandleFunc("/result", controller.ResultPage)
	http.HandleFunc("/restart", controller.Restart)
}
