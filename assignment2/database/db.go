package database

import (
	"log"

	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() (*gorm.DB, error) {

	conn := ""

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error saat terhubung ke database", err)
	}
	//fmt.Println("Database has been connected")

	db.AutoMigrate(&models.Item{}, &models.Order{})

	return db, nil
}
