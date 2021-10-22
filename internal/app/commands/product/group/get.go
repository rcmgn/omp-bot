package group

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ProductGroupCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	group, err := c.groupService.Describe(id)
	if err != nil {
		log.Printf("fail to get group id %d: %v", id, err)
		return
	}
	log.Printf("Get  group =  %v  msg= %s", group, group.Owner+" "+group.Items)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		group.Owner+" "+group.Items,
	)

	c.bot.Send(msg)
}
