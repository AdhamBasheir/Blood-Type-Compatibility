package main

import (
	"blood-type-compatibility/initializers"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	Migrate()
}
