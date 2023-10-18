package database

import (
	"fmt"
	"github.com/uzzalhcse/GoSpire/app/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Post{},
		&models.User{},
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
