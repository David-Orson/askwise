package models

type User struct {
	Base
	GoogleID  string `gorm:"uniqueIndex"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	ImageURL  string
	Projects  []Project
	Documents []Document
}
