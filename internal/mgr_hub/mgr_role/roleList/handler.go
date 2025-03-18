package roleList

func (x *api) handler() (error, *RoleList_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err, roles := x.service()
	if err != nil {
		return err, nil
	}
	return nil, &RoleList_Output{
		ListRole: roles,
	}
}
