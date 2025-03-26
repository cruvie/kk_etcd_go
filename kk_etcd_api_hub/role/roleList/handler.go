package roleList

func (x *api) Handler() (*RoleList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	roles, err := x.service()
	if err != nil {
		return nil, err
	}
	return &RoleList_Output{
		ListRole: roles,
	}, nil
}
