package group

import (
	"fmt"
	"strings"

	"github.com/rcmgn/omp-bot/internal/model/product"
)

func (c *DummyGroupService) Create(group product.Group) (int, error) {
	group.Owner = strings.Trim(group.Owner, " ")

	if group.Owner == "" {
		return 0, fmt.Errorf("group owner is empty")
	}
	group.Items = strings.Trim(group.Items, " ")

	if group.Items == "" {
		return 0, fmt.Errorf("group items is empty")
	}

	newID := len(Groups) + 1

	Groups[newID] = product.Group{
		ID:    newID,
		Owner: group.Owner,
		Items: group.Items,
	}

	return newID, nil
}
