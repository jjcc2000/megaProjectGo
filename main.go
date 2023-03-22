package main

import (

	"log"
	"net/http"
	"path/mod/db"
	"path/mod/models"
	"path/mod/routes"

	"github.com/gorilla/mux"
)

func main() {
	db.CheckDB()
	db.DB.AutoMigrate(models.Users{})


	r := mux.NewRouter()

	r.HandleFunc("/getForm",routes.SubmitFormHandler)
	r.HandleFunc("/submit-form",routes.FormGetter)
	r.HandleFunc("/", routes.GetUsersHandles).Methods("GET")
	r.HandleFunc("/user/{id}",routes.GetUser).Methods("GET")
	r.HandleFunc("/", routes.CreateUsers).Methods("POST")
	r.HandleFunc("/user/{id}",routes.DeleteUsers).Methods("DELETE")

		
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", r))	
}
