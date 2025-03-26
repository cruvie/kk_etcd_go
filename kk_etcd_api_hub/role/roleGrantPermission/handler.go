package roleGrantPermission

func (x *api) Handler() (*RoleGrantPermission_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return nil, err
	}
	return &RoleGrantPermission_Output{}, nil
}
