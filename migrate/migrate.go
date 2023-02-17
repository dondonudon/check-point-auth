package main

import (
	"fmt"
	"log"

	"github.com/dondonudon/authentication-authorization/initializers"
	"github.com/dondonudon/authentication-authorization/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("Migration complete")
}
