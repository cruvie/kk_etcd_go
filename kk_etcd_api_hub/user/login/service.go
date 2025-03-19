package login

func (x *api) service() (tokenString string, err error) {
	span := x.stage.StartTrace("service")
	defer span.End()

	//todo gen jwt
	return "todo gen jwt from etcd", nil
}
