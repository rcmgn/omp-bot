package group

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ProductGroupCommander) Start(inputMsg *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		"Start the bot",
	)

	c.bot.Send(msg)
}
