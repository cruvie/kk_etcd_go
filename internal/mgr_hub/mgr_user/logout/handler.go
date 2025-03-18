package logout

func (x *api) handler() (error, *Logout_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return err, &Logout_Output{}
}
