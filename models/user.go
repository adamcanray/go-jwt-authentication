package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// to pluralize table name, see: https://gorm.io/docs/conventions.html#Pluralized-Table-Name
type User struct {
	gorm.Model
	Name     string `json:"name"`                   // reflect
	Username string `json:"username" gorm:"unique"` // duplicate value err will handled by GORM
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// the reason we use pointer of *User is we want assign Password of user by specifics memory address?
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
