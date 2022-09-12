package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahmaabdelaziiz/webApp/data"
)

type SignupRequest struct {
	Name     string
	Email    string
	Password string
}

type SigninRequest struct {
	// Name     string
	Email    string
	Password string
}

func main() {
	app := fiber.New()

	engine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	app.Post("/signup", func(c *fiber.Ctx) error {
		req := new(SignupRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		if req.Name == "" || req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid SignUp Credentials")
		}

		return nil

	})
	app.Post("/signin", func(c *fiber.Ctx) error {
		req := new(SigninRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		return nil

	})
	app.Get("/private", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path": "private"})

	})
	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path": "public"})

	})
	(app.Listen(":8080"))
}
