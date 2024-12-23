package main

import (
	"log"

	"github.com/joho/godotenv"

	"gopkg.in/telebot.v4"
	telemiddleware "gopkg.in/telebot.v4/middleware"

	"github.com/xoticdsign/pingpongerbot/internal/app"
	"github.com/xoticdsign/pingpongerbot/internal/handlers"
	"github.com/xoticdsign/pingpongerbot/internal/middleware"
)

func main() {
	godotenv.Load()

	bot, err := app.InitApp(false)
	if err != nil {
		log.Fatalf(" ERR %v", err)
	}

	bot.Use(telemiddleware.Recover())
	bot.Use(middleware.Log)

	unrecognized := bot.Group()

	unrecognized.Handle(telebot.OnText, handlers.Start)
	unrecognized.Handle(telebot.OnMedia, handlers.Start)

	bot.Handle("/start", handlers.Start)
	bot.Handle("/ping", handlers.Ping)
	bot.Handle("/echo", handlers.Echo)

	bot.Start()
}
