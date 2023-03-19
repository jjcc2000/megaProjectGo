package models

import "gorm.io/gorm"

type Users struct {
	//Especified that it is a gorm.Models
	gorm.Model
	FirstName string
	LastName  string
	Age       int
}

