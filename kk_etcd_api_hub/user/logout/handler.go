package logout

func (x *api) Handler() (*Logout_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return &Logout_Output{}, err
}
