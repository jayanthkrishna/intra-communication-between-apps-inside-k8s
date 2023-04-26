package main

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/api", randomnumberapi)

	app.Listen(":4040")

}

func randomnumberapi(c *fiber.Ctx) error {

	r := "Generating Random int " + fmt.Sprint(rand.Intn(100000))
	return c.SendString(r)
}
