package main

import (
	"fmt"
	"log"

	goenv "github.com/Netflix/go-env"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Env struct {
	Port string `env:"PORT,required,default=3000"`
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	var env Env
	_, err := goenv.UnmarshalFromEnviron(&env)
	if err != nil {
		log.Fatalf("Could not load environment variables: %s", err)
	}

	app.Get("/", func(c fiber.Ctx) error {
		fmt.Println(c.Host())
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + env.Port)
}
