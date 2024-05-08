package staffHandler

import (
	// "fmt"
	// "strings"
	// "time"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
		"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Hakimaris/personalProject_eniqilo/helper"
	"github.com/Hakimaris/personalProject_eniqilo/model"

)

func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}

func Login(c *fiber.Ctx) error {
	return c.SendString("Register")
}
