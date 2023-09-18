package cmd

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string   `json`
	PhoneNumber []string `json`
	Email       string   `json`
}

func (s *User) tableName() string {
	return "users"
}
