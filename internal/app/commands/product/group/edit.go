package group

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rcmgn/omp-bot/internal/model/product"
)

func (c *ProductGroupCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	id := strings.Split(args, " ")[0]

	groupID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("failed to parse id: %v", err)
		return
	}

	group := product.Group{
		ID:    groupID,
		Owner: strings.Split(args, " ")[1],
		Items: strings.Split(args, " ")[2],
	}

	err = c.groupService.Update(groupID, group)
	if err != nil {
		log.Printf("fail to edit comment with id %d: %v", groupID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Group %d edited", groupID),
	)

	c.bot.Send(msg)
}
