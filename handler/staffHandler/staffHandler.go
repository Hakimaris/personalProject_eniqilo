package staffHandler

import (
	// "fmt"
	// "strings"
	// "time"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
)

func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}

func Login(c *fiber.Ctx) error {
	return c.SendString("Register")
}
