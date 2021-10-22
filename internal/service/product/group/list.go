package group

import (
	"github.com/rcmgn/omp-bot/internal/model/product"
)

func (c *DummyGroupService) List() ([]product.Group, error) {
	return c.groups, nil
}
