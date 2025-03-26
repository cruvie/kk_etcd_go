package userList

func (x *api) Handler() (*UserList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	users, err := x.service()
	if err != nil {
		return nil, err
	}
	return &UserList_Output{
		ListUser: users,
	}, err
}
