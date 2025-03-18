package roleGrantPermission

func (x *api) handler() (error, *RoleGrantPermission_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return err, nil
	}
	return nil, &RoleGrantPermission_Output{}
}
