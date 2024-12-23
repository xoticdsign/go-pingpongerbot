package handlers

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v4"
)

func Start(c telebot.Context) error {
	greeting := "Привет, меня зовут pingponger! 🚀 Выбери одну из комманд ниже, чтобы узнать, как я работаю."

	toSend := []string{
		greeting,
	}

	var cmdsSlice []string

	cmds, _ := c.Bot().Commands()

	for _, cmd := range cmds {
		if cmd.Text == "start" {
			continue
		}
		cmdStr := fmt.Sprintf("/%v - %v", cmd.Text, cmd.Description)
		cmdsSlice = append(cmdsSlice, cmdStr)
	}

	cmdsList := strings.Join(cmdsSlice, "\n")
	toSend = append(toSend, cmdsList)
	toSendFmt := strings.Join(toSend, "\n\n")

	return c.Send(toSendFmt)
}

func Ping(c telebot.Context) error {
	return c.Send("pong!")
}

func Echo(c telebot.Context) error {
	args := c.Args()

	if len(args) == 0 {
		return c.Send("Используй - /echo <твое сообщение здесь>")
	}

	arg := strings.Join(args, " ")

	return c.Reply(arg)
}
