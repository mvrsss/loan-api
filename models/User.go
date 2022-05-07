package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID         string  `gorm:"column:uid type:text"`
	Name        string  `gorm:"column:name" json:"name"`
	Surname     string  `gorm:"column:surname" json:"surname"`
	IIN         string  `gorm:"column:iin" json:"iin"`
	PhoneNumber string  `gorm:"column:phonenumber" json:"phonenumber"`
	Password    string  `gorm:"column:password" json:"password"`
	Address     string  `gorm:"column:address" json:"address"`
	Balance     float64 `gorm:"column:balance"`
	Token       string  `gorm:"column:token"`
}

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
	// bb, _ := bcrypt.GenerateFromPassword([]byte(providedPassword), 14)
	// fmt.Println(string(bb))
	// fmt.Println(user.Password)
	if err != nil {
		return err
	}
	return nil
}
