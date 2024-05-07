// package main

// import (
// 	"fmt"
// 	"strings"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func authRequired(ctx *fiber.Ctx) error {
// 	authHeader := ctx.Get("Authorization")
// 	if authHeader == "" {
// 		return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization Header")
// 	}

// 	parts := strings.Split(authHeader, " ")
// 	if !(len(parts) == 2 && parts[0] == "Bearer") {
// 		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization Header")
// 	}

// 	tokenString := parts[1]
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// Return the key for validation
// 		return []byte("secret"), nil
// 	})

// 	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
// 		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Token")
// 	}

// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Token")
// 	}

// 	ctx.Locals("user", token)
// 	return ctx.Next()
// }

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(ctx *fiber.Ctx) error {
// 		ctx.Send([]byte("Hello, World!"))
// 		return nil
// 	})

// 	app.Post("/login", Login)

// 	app.Post("/hello", authRequired, func(ctx *fiber.Ctx) error {
// 		ctx.Send([]byte("Hello, Protect!"))
// 		return nil
// 	})

// 	err := app.Listen(":8080") // Convert the port number to a string
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func Login(ctx *fiber.Ctx) error {
// 	type request struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	var body request
// 	err := ctx.BodyParser(&body)
// 	if err != nil {
// 		ctx.Status(fiber.StatusBadRequest).JSON(
// 			fiber.Map{
// 				"error": "CANNOT PARSE JSON",
// 			})
// 		return nil
// 	}

// 	if body.Email != "body@gmail.com" || body.Password != "123456" {
// 		ctx.Status(fiber.StatusUnauthorized).JSON(
// 			fiber.Map{
// 				"error": "INVALID CREDENTIALS",
// 			})
// 		return nil
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
// 	claims["sub"] = "1"

// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		ctx.Status(fiber.StatusInternalServerError).JSON(
// 			fiber.Map{
// 				"error": "CANNOT CREATE TOKEN",
// 			})
// 		return nil
// 	}

// 	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"token": t,
// 		"user": struct {
// 			Id    int    `json:"id"`
// 			Email string `json:"email"`
// 		}{
// 			Id:    1,
// 			Email: body.Email,
// 		},
// 	})
// 	return nil
// }

