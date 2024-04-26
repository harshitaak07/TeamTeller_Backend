package models

import (
	"gorm.io/gorm"
)

// Story is a representation of the story model
type Story struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null;" json:"title" valid:"required~Title is required"`
	Description string `gorm:"type:text;not null;" json:"description" valid:"required~Description is required"`
	Content     string `gorm:"type:text;not null;" json:"content" valid:"required~Content is required"`
	AuthorID    uint   `gorm:"not null;" json:"author_id" valid:"required~Author ID is required"`
	Author      User   `gorm:"foreignKey:AuthorID" json:"author"`
	Genre       string `gorm:"type:varchar(100);not null;" json:"genre" valid:"required~Genre is required"`
	LikeCount   int    `gorm:"type:integer;not null;" json:"like_count"`
}
