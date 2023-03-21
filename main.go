package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"path/mod/db"
	"path/mod/models"
	"path/mod/routes"
)

func main() {
	db.CheckDB()
	db.DB.AutoMigrate(models.Users{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.GetUsersHandles).Methods("GET")
	r.HandleFunc("/user/{id}",routes.GetUser).Methods("GET")
	r.HandleFunc("/", routes.CreateUsers).Methods("POST")
	r.HandleFunc("/user/{id}",routes.DeleteUsers).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
