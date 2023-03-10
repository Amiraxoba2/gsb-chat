package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, _ = gorm.Open(sqlite.Open("resource/gsb.sqlite"), &gorm.Config{})
