package myInfo

func (x *api) service() error {
	span := x.stage.StartTrace("service")
	defer span.End()

	return nil
}
