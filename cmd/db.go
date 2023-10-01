package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Phone string
}

var models = []interface{}{
	&User{},
}
var DB *gorm.DB

func InitDB() {
	dsn := "host=127.0.0.1 user=postgres password=tachtach12203 dbname=gorm port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connection successfully opened.")
	for _, model := range models {
		err := DB.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Could not auto migrate model %T: %v", model, err)
		}
	}
	log.Println("Schema successfully auto-migrated.")
}
