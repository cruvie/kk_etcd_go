package roleDelete

import (
	"context"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *Api) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("service")
	defer span.End()

	_, err := global_model.GetClient(stage).RoleDelete(context.Background(), x.In.GetName())
	if err != nil {
		return err
	}
	return nil
}
