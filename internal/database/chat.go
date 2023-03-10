package database

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Name     string
	Messages string
}
