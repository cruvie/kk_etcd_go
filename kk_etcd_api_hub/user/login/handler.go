package login

func (x *api) Handler() (*Login_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	tokenString, err := x.service()
	return &Login_Output{
		Token: tokenString,
	}, err
}
