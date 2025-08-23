package login

import "gitee.com/cruvie/kk_go_kit/kk_stage"

func (x *Api) Service(stage *kk_stage.Stage) (tokenString string, err error) {
	span := stage.StartTrace("service")
	defer span.End()

	// todo gen jwt
	return "todo gen jwt from etcd", nil
}
