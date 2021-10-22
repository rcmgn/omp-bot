package group

import "github.com/rcmgn/omp-bot/internal/model/product"

var Groups = map[int]product.Group{
	1: {
		ID:    1,
		Owner: "Owner 1",
		Items: "Item 1",
	},
	2: {
		ID:    2,
		Owner: "Owner 1",
		Items: "Item 2",
	},
	3: {
		ID:    3,
		Owner: "Owner 1",
		Items: "Item 3",
	},
	4: {
		ID:    4,
		Owner: "Owner 2",
		Items: "Item 1",
	},
	5: {
		ID:    5,
		Owner: "Owner 2",
		Items: "Item 2",
	},
	6: {
		ID:    6,
		Owner: "Owner 2",
		Items: "Item 3",
	},
}

type Subdomain struct {
	Title string
}
