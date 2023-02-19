package models

import "gorm.io/gorm"

// create a todo model
type Todo struct {
	gorm.Model
	Name string `json:"name"`
}
