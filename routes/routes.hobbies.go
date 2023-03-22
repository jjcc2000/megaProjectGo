package routes

import (
	"html/template"
	"net/http"
)

var TmplHobbies *template.Template
func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	TmplHobbies= template.Must(template.ParseFiles("form.html"))
	TmplHobbies.ExecuteTemplate(w,"form.html",nil)
}
