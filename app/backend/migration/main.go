package main

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/models"

	"github.com/sirupsen/logrus"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func Migrate() {
	err := initializers.DB.AutoMigrate(&models.BloodType{}, &models.User{})
	if err != nil {
		logrus.WithError(err).Fatal("Database migration failed")
	}
	logrus.Info("Database migration completed successfully")
}

func main() {
	Migrate()
}
