package group

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rcmgn/omp-bot/internal/model/product"
)

func (c *ProductGroupCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	texts := strings.Split(args, ` `)

	group := product.Group{
		Owner: texts[0],
		Items: texts[1],
	}

	id, err := c.groupService.Create(group)
	if err != nil {
		log.Printf("fail to create group: %v", err)
	}

	response := fmt.Sprintf("Group with %v id's created", id)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		response,
	)

	c.bot.Send(msg)
}
