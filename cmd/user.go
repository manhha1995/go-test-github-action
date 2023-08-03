package cmd

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `json`
}

func (s *User) tableName() string {
	return "users"
}
