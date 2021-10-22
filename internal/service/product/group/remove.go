package group

func (c *DummyGroupService) Remove(group_id int) (bool, error) {
	delete(Groups, group_id)
	return true, nil
}
