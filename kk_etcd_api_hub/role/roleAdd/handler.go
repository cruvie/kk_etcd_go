package roleAdd

func (x *api) Handler() (*RoleAdd_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return nil, err
	}
	return &RoleAdd_Output{}, nil
}
