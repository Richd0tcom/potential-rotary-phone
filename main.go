package main

import (
	"log"
	"os"

	"github.com/Richd0tcom/potential-rotary-phone/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "1738"
	}

	app:= fiber.New()

	app.Server().StreamRequestBody = true

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	

	app.Get("/healthcheck", func(fiberContext *fiber.Ctx) error {
		return fiberContext.SendString("OK")
	})

	app.Post("/api/pre-upload", api.HandlePreUpload)
	app.Post("/api/upload", api.HandleUpload)
	app.Get("/api/video-details", api.ServeVideoData)
	app.Get("/api/docs", api.RedirectToDocs)

	log.Fatal(app.Listen(`:`+ port))
}