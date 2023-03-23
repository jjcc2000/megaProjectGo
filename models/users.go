package models

import "gorm.io/gorm"

type Users struct {
	//Especified that it is a gorm.Models
	gorm.Model
	Id int
	FirstName string
	LastName  string
	Age       int
	Hobbies []Hobbies
}

