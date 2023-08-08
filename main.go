package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"bitly/internal/app"
)

func main() {
	app.SetupDB()   // start database

	router := fiber.New()   // new fiber
	router.Use(cors.New(cors.Config{   // cors
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/goly", app.GetGollies)
	router.Get("/goly/:id", app.GetGoly)
	router.Post("/goly", app.CreateGoly)
	router.Patch("/goly", app.UpdateGoly)
	router.Delete("goly/:id", app.DeleteGoly)
	
	router.Get("/r/:redirect", app.Redirect)
	
	router.Listen(":3000")
}