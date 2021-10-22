package group

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rcmgn/omp-bot/internal/service/product/group"
)

func (c *ProductGroupCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText := ""
	for id := range group.Groups {
		outputMsgText += fmt.Sprintf("%d. %v \n", id, group.Groups[id])
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	c.bot.Send(msg)
}
