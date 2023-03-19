package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/mod/db"
	"path/mod/models"
)

var tmpl *template.Template

func GetUsersHandles(w http.ResponseWriter, r *http.Request) {
	//Need an interface to interact with the database
	var Dt []models.Users
	//Looks for this struct in to the database
	db.DB.Find(&Dt)
	//Need a way to Template the HTML
	tmpl = template.Must(template.ParseFiles("index.html"))
	//Struct to test if the template works
	tmpl.Execute(w,&Dt)
}

func CreateUsers(w http.ResponseWriter, r *http.Request){
	var recievedUsers models.Users
	json.NewDecoder(r.Body).Decode(&recievedUsers)
	//Use the method Create to Create new Table with the request 
	checker := db.DB.Create(&recievedUsers)
	err:= checker.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Create Users Failed"))
		return 
	}
	json.NewEncoder(w).Encode(&recievedUsers)
}