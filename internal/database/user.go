package database

import "gorm.io/gorm"

type Role int

const (
	Admin Role = iota + 1
	Mod
	Member
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Role
}
