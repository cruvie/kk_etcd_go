package login

func (x *api) handler() (error, *Login_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	tokenString, err := x.service()
	return err, &Login_Output{
		Token: tokenString,
	}
}
