package roleAdd

func (x *api) handler() (error, *RoleAdd_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return err, nil
	}
	return nil, &RoleAdd_Output{}
}
