package main

import (
	"fmt"
	"hedwig/cmd/smtp"
	"hedwig/config"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	f *fiber.App
	c *config.Config
}

func NewServer(config *config.Config) (server *Server) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/hello/:to/:title/:message/:img?", func(ctx *fiber.Ctx) error {
		to := ctx.Params("to")
		title := ctx.Params("title")
		message := ctx.Params("message")
		image := ctx.Params("img")

		if len(image) <= 0 {
			image = "https://media.giphy.com/media/dLPHtoQY2l7glPL1qk/giphy.gif"
		}

		mailCx := smtp.BumMail{Smtp: config.Smtp}

		err := mailCx.Send([]string{to}, title, message, image)

		if err != nil {
			return ctx.SendString("Something went wrong 🧨")
		} else {
			return ctx.SendString("Email Sent 🚀")
		}
	})

	app.Post("/hello", func(ctx *fiber.Ctx) error {
		params := new(smtp.EmailRequest)

		err := ctx.BodyParser(params)

		if err != nil {
			ctx.Response().SetStatusCode(409)
			return ctx.SendString("Something went wrong 🧨")
		}

		if len(params.Image) <= 0 {
			params.Image = "https://media.giphy.com/media/dLPHtoQY2l7glPL1qk/giphy.gif"
		}

		mailCx := smtp.BumMail{Smtp: config.Smtp}

		err = mailCx.Send([]string{params.To}, params.Title, params.Message, params.Image)

		if err != nil {
			ctx.Response().SetStatusCode(409)
			return ctx.SendString("Something went wrong 🧨")
		} else {
			return ctx.SendString("Email Sent 🚀")
		}
	})

	server = &Server{
		f: app,
		c: config,
	}

	return server
}

func (s *Server) Start() error {
	return s.f.Listen(fmt.Sprintf(":%d", s.c.Server.Port))
}
