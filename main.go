package main

import (
	"fmt"
	"net/http"
	"os"

	"rest-api/app"
	"rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/operation/{id}", controllers.Transaction).Methods("POST")
	router.HandleFunc("/api/operation/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/operation/date/{operation_date}", controllers.GetDate).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
