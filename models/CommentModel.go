package models

import (
	"gorm.io/gorm"
)

// Comment is a representation of the comment model
type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"not null;" json:"user_id" valid:"required~User ID is required"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	StoryID uint   `gorm:"not null;" json:"story_id" valid:"required~Story ID is required"`
	Story   Story  `gorm:"foreignKey:StoryID" json:"story"`
	Content string `gorm:"type:text;not null;" json:"content" valid:"required~Content is required"`
}
