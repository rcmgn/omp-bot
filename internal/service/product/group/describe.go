package group

import (
	"fmt"

	"github.com/rcmgn/omp-bot/internal/model/product"
)

func (c *DummyGroupService) Describe(group_id int) (*product.Group, error) {
	if len(Groups) <= group_id {
		return nil, fmt.Errorf("id does not exist")
	}
	group := Groups[group_id]

	return &group, nil
}
