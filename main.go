package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/trioxtron/golang-service/api"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/", func(c *fiber.Ctx) error {
        return api.GetApis(c)
	})
	app.Get("/api/:api", func(c *fiber.Ctx) error {
        return api.GetApi(c)
	})
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")

}
