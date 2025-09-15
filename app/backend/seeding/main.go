package main

import (
	"blood-type-compatibility/initializers"
	"os"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	SeedBloodTypes()
	if os.Getenv("APP_ENV") == "dev" {
		SeedTestUsers()
	}
}
