package group

import (
	"fmt"
	"strings"

	"github.com/rcmgn/omp-bot/internal/model/product"
)

func (c *DummyGroupService) Update(group_id int, group product.Group) error {
	group.Owner = strings.Trim(group.Owner, " ")

	if group.Owner == "" {
		return fmt.Errorf("Group's owner is empty")
	}

	Groups[group_id] = group

	return nil
}
