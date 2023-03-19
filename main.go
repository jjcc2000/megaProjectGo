package main

import (
	"net/http"
	"path/mod/db"
	"path/mod/models"
	"path/mod/routes"
	"github.com/gorilla/mux"
)

func main(){
	db.CheckDB()
	db.DB.AutoMigrate(models.Users{})

	r := mux.NewRouter()
	r.HandleFunc("/",routes.GetUsersHandles).Methods("GET")
	r.HandleFunc("/",routes.CreateUsers).Methods("POST")

	http.ListenAndServe(":8080",r)
}
