package models

import (
	"gorm.io/gorm"
)

// Like is a representation of the like model
type Like struct {
	gorm.Model
	UserID  uint  `gorm:"not null;" json:"user_id" valid:"required~User ID is required"`
	User    User  `gorm:"foreignKey:UserID" json:"user"`
	StoryID uint  `gorm:"not null;" json:"story_id" valid:"required~Story ID is required"`
	Story   Story `gorm:"foreignKey:StoryID" json:"story"`
}
