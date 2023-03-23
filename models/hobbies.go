package models

import "gorm.io/gorm"

type Hobbies struct{
	//To specified that this is a gorm Model
	gorm.Model
	
	Hobbies string 
	Description string
	Hours int 
	UsersId uint

}