package main

import (
	// "fmt"
	// "strings"
	// "time"
	staffHandler "github.com/Hakimaris/personalProject_eniqilo"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	api := app.Group("/api") // /api

	v1 := api.Group("/v1")                 // /api/v1
	v1.Get("/list", staffHandler.register) // /api/v1/list
	v1.Post("/user", staffHandler.login)   // /api/v1/user

	// v2 := api.Group("/v2")        // /api/v2
	// v2.Get("/list", handler)      // /api/v2/list
	// v2.Get("/user", handler)      // /api/v2/user

	log.Fatal(app.Listen(":8080"))
}
