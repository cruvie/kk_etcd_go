package roleRevokePermission

func (x *api) Handler() (*RoleRevokePermission_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return nil, err
	}
	return &RoleRevokePermission_Output{}, nil
}
