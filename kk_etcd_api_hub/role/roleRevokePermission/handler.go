package roleRevokePermission

func (x *api) Handler() (error, *RoleRevokePermission_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return err, nil
	}
	return nil, &RoleRevokePermission_Output{}
}
