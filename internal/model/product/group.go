package product

import "fmt"

type Group struct {
	ID    int
	Owner string
	Items string
}

func (c *Group) String() string {
	return fmt.Sprintf("ID: %d, Owner: %s, Items: %s\n", c.ID, c.Owner, c.Items)
}
