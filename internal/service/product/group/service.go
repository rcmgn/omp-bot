package group

import (
	"github.com/rcmgn/omp-bot/internal/model/product"
)

type GroupService interface {
	Describe(groupID int) (*product.Group, error)
	List() ([]product.Group, error)
	Create(product.Group) (int, error)
	Update(groupID int, group product.Group) error
	Remove(groupID int) (bool, error)
}

type DummyGroupService struct {
	groups []product.Group
}

func NewDummyGroupService() *DummyGroupService {
	return &DummyGroupService{}
}
