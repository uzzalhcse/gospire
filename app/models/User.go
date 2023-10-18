package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Phone    string `gorm:"size:255;not null;unique" json:"phone"`
	Password string `gorm:"size:255;not null;" json:"password"`
	TimeStamps
}

func (User) TableName() string {
	return "users" // Replace with your actual table name
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) FormatResponse() interface{} {
	return map[string]interface{}{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"phone": u.Phone,
	}
}
