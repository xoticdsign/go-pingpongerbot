package app

import (
	"os"
	"time"

	"gopkg.in/telebot.v4"

	"github.com/xoticdsign/pingpongerbot/internal/middleware"
)

func InitApp(offline bool) (*telebot.Bot, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token: os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{
			Limit:   50,
			Timeout: time.Second * 15,
		},
		OnError: middleware.Error,
		Offline: offline,
	})
	if err != nil {
		return nil, err
	}
	return bot, nil
}
