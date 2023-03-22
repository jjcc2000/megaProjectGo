package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"path/mod/db"
	"path/mod/models"
	"strconv"
)

var TmplHobbies *template.Template
var TmplThanks *template.Template

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	TmplHobbies= template.Must(template.ParseFiles("form.html"))
	TmplHobbies.ExecuteTemplate(w,"form.html",nil)
}


func FormGetter(w http.ResponseWriter, r *http.Request){
	var Users  models.Users
	//The key is the name in the Html object
	Users.FirstName = r.FormValue("fName")
	Users.LastName = r.FormValue("lName")	  
	//Parse the in to a string to save in the database 
	stri , err:=strconv.Atoi(r.FormValue("Age"))
	if err != nil {
		fmt.Fprintln(w,"The Age was not an integer")
		return 
	}
	Users.Age = stri
	//Save in the database
	fmt.Println("Method has been called")
	checker := db.DB.Create(&Users)
	//Check if there was an error while saving in database
	err=checker.Error
	if err != nil {
		fmt.Fprintln(w,"There was an error while saving in the database")
		return
	}
	TmplThanks = template.Must(template.ParseFiles("thanks.html"))
	TmplThanks.ExecuteTemplate(w,"thanks.html",nil)

}
