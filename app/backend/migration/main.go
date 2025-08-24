package main

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"
	"log"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func Migrate() {
	err := initializers.DB.AutoMigrate(&models.BloodType{}, &models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed successfully")
}

func main() {
	Migrate()
}
