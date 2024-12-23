package middleware

import (
	"strings"

	"gopkg.in/telebot.v4"
)

func Error(err error, c telebot.Context) {
	msg := "Ошибочка :<"
	description := "Поделись ошибкой с @xoticdsign."

	toSend := []string{
		msg,
		description,
		err.Error(),
	}

	toSendFmt := strings.Join(toSend, "\n\n")

	c.Send(toSendFmt)
}
