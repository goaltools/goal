package controllers_test

import (
	"log"

	"github.com/anonx/sunplate/internal/skeleton/assets/handlers"

	c "github.com/anonx/sunplate/config"
)

func init() {
	// Openning configuration file.
	config := c.New()
	err := config.ParseFile("config/config.ini", c.ReadFromFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialization of handlers.
	handlers.Init(config)
}
