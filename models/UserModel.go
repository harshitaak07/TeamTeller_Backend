package models

import (
	"gorm.io/gorm"
)

// User is a representation of the user model
type User struct {
	gorm.Model
	Email      string `gorm:"type:varchar(100);not null;uniqueIndex" json:"email" valid:"required~Email is required,email~Email is not valid"`
	Name       string `gorm:"type:varchar(100);not null;" json:"name" valid:"required~Name is required,matches(^[a-zA-Z ]+$)~Name must be alphabetic"`
	ProfilePic string `gorm:"type:text;not null;" json:"profile_pic" valid:"required~Profile picture is required"`
	YOB        int    `gorm:"type:smallint;not null;" json:"yob" valid:"required~Year of birth is required,range(1900|2005)~Year of birth must be between 1900 and 2005"`
}
