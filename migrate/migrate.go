package main

import (
	"fmt"
	"log"

	"github.com/W5DEV/go-recipe-api/initializers"
	"github.com/W5DEV/go-recipe-api/models"
)


func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Recipe{})
	fmt.Println("? Migration complete")
}

