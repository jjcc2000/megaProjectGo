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


//Function to call the form html
func SubmitForm(w http.ResponseWriter, r *http.Request) {
	TmplHobbies= template.Must(template.ParseFiles("form.html"))
	TmplHobbies.ExecuteTemplate(w,"form.html",nil)
}

//Function to get The values of submit-form and call the thanks form
func FormAddUssers(w http.ResponseWriter, r *http.Request){
	fmt.Println("The Form Getter Ussers")
	var Users  models.Users
	//The key is the name in the Html object
	Users.FirstName = r.FormValue("fName")
	Users.LastName = r.FormValue("lName")	  
	//Parse the in to a string to save in the database 
	stri , err:=strconv.Atoi(r.FormValue("Age"))
	if err != nil {
		fmt.Println("Fail to convert the age to string")
		fmt.Fprintln(w,"The Age was not an integer")
		return 
	}
	Users.Age = stri
	//Save in the database
	fmt.Println("Form Getter Method has been called")
	checker := db.DB.Create(&Users)
	//Check if there was an error while saving in database
	err=checker.Error
	if err != nil {
		fmt.Println("Error in the Creation  of the database")
		fmt.Fprintln(w,"There was an error while saving in the database")
		return
	}
	TmplHobbies = template.Must(template.ParseFiles("thanks.html"))
	TmplHobbies.Execute(w,&Users)
}
//Function that is called to deleted users
func DeletedUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println("The Form Deleted Users")
	var UserToDeleted  models.Users
	idStr := r.FormValue("idD")
	db.DB.First(&UserToDeleted,idStr)
	if UserToDeleted.ID==0{
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("No user has that Id")
		fmt.Fprintln(w,"There is no match for that id")
		return
	}
	db.DB.Unscoped().Delete(&UserToDeleted)
	fmt.Println("The user has been deleted")
	fmt.Fprintln(w,"The user has been deleted")
	w.WriteHeader(http.StatusOK)
}



