package logout

import "gitee.com/cruvie/kk_go_kit/kk_stage"

func (x *Api) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("service")
	defer span.End()

	// todo remove jwt from etcd
	return nil
}
