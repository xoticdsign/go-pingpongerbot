package middleware

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/telebot.v4"
)

func initLogger() zerolog.Logger {
	logger := zerolog.New(&zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeFormat:   "2006-01-02 | 15:04:05.000",
		TimeLocation: time.UTC,
	}).With().Timestamp().Logger()

	return logger
}

func Log(next telebot.HandlerFunc) telebot.HandlerFunc {
	logger := initLogger()

	return func(c telebot.Context) error {
		logger.Info().
			Int64("ID", c.Chat().ID).
			Str("USERNAME", c.Sender().Username).
			Str("TEXT", c.Text()).
			Msg("UPDATE")

		err := next(c)
		if err != nil {
			logger.Error().
				Int64("ID", c.Chat().ID).
				Str("USERNAME", c.Sender().Username).
				Str("TEXT", c.Text()).
				Msg(err.Error())
		}

		return err
	}
}
