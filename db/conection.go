package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//This is a pointer to a gorm DB object
var DB  *gorm.DB

//Define the coneecction to the database and its credentials
//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

var stdb = "root:sqlpassword@tcp(127.0.0.1:3306)/gotemplate?charset=utf8mb4&parseTime=True&loc=Local"
//Check if the connection to the database is working
func CheckDB(){
	var err error
	DB, err = gorm.Open(mysql.Open(stdb),&gorm.Config{})
	if err != nil {
		log.Fatal("There was an error conecting to the database")
	}else{
		log.Println("Conected to the database")
	}
}