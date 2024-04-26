package models

import (
	"gorm.io/gorm"
)

// Collaboration is a representation of the collaboration model
type Collaboration struct {
	gorm.Model
	StoryID uint   `gorm:"not null;" json:"story_id" valid:"required~Story ID is required"`
	Story   Story  `gorm:"foreignKey:StoryID" json:"story"`
	UserID  uint   `gorm:"not null;" json:"user_id" valid:"required~User ID is required"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	Status  string `gorm:"type:varchar(100);not null;" json:"status" valid:"required~Status is required"`
}
