package controllers_test

import (
	"log"

	"github.com/colegion/goal/internal/skeleton/assets/handlers"

	c "github.com/colegion/goal/config"
)

func init() {
	// Openning configuration file.
	config := c.New()
	err := config.ParseFile("../config/testing.ini", c.ReadFromFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialization of handlers.
	handlers.Init(config)
}
