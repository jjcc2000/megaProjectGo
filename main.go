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
	db.DB.AutoMigrate(models.Hobbies{})

	r := mux.NewRouter()
	//Calls the first page In which you selected stuff
	r.HandleFunc("/", routes.FirstPage)
	//Servers the form html
	r.HandleFunc("/form.html",routes.SubmitForm)
	//When the submit-form is called 
	r.HandleFunc("/submit-form",routes.FormAddUssers).Methods("GET")
	//When the deleted-form is called
	r.HandleFunc("/delete-form",routes.DeletedUsers).Methods("GET")
	//When the edite Form is called
	r.HandleFunc("/update-form",routes.ThankHandlers)
	
	r.HandleFunc("/user/{id}",routes.GetUser).Methods("GET")
	r.HandleFunc("/", routes.CreateUsers).Methods("POST")
	r.HandleFunc("/user/{id}",routes.DeleteUsers).Methods("DELETE")

		
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", r))	
}
