package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/mod/db"
	"path/mod/models"
	"github.com/gorilla/mux"
)

var tmpl *template.Template


func GetUsersHandles(w http.ResponseWriter, r *http.Request) {
	//Need an interface to interact with the database
	var Dt []models.Users
	//Looks for this struct in to the database
	db.DB.Find(&Dt)
	//Need a way to Template the HTML
	tmpl = template.Must(template.ParseFiles("index.html"))
	//Executes the template
	err:=tmpl.Execute(w,&Dt)
	if err != nil {
		fmt.Println("There has been an error in the template creations",err)
	}
}
func GetUser(w http.ResponseWriter, r *http.Request){
	var Dt []models.Users
	vars := mux.Vars(r)
	db.DB.First(&Dt,vars["id"])
	if len(Dt)==0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("That id does not Exist"))
		return
	}
	tmpl = template.Must(template.ParseFiles("index.html"))
	err:= tmpl.Execute(w,&Dt)
	if err != nil {
		fmt.Println("There is an error in the GetUser Method",err)
	}
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
}
func DeleteUsers(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var user  models.Users
	json.NewDecoder(r.Body).Decode(&user)

	db.DB.First(&user,vars["id"])
	if user.ID==0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("That id does not match any registered"))
	}else{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The user has been deleted"))
		db.DB.Unscoped().Delete(&user)
	}
}	
