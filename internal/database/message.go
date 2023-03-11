package database

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Chat    uint
	Content string
	Author  uint
}
