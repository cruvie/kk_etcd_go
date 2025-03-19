package roleGet

func (x *api) service() (err error) {
	span := x.stage.StartTrace("service")
	defer span.End()

	return nil
}
