package userList

func (x *api) Handler() (error, *UserList_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err, users := x.service()
	if err != nil {
		return err, nil
	}
	return err, &UserList_Output{
		ListUser: users,
	}
}
