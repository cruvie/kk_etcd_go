package logout

func (x *api) service() error {
	span := x.stage.StartTrace("service")
	defer span.End()

	//todo remove jwt from etcd
	return nil
}
