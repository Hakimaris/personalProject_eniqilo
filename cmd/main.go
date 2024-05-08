package main

import (
	// "fmt"
	// "strings"
	// "time"
	staffHandler "github.com/Hakimaris/personalProject_eniqilo/handler/staffHandler"
	// customerHandler "github.com/Hakimaris/personalProject_eniqilo/handler/customerHandler"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")
	
	// product := v1.Group("/product") // /api/v1/product
	
	
	staff := v1.Group("/staff")           
	staff.Post("/register", staffHandler.Register) 
	staff.Post("/login", staffHandler.Login)   

	// customer := v1.Group("/customer")      // /api/v1/customer
	// customer.Get("/register", customerHandler.Register)
	// customer.Post("/login", customerHandler.Login)


	log.Fatal(app.Listen(":8080"))
}
