package userDelete

import (
	"context"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *Api) Service(stage *kk_stage.Stage, userName string) error {
	span := stage.StartTrace("service")
	defer span.End()

	_, err := global_model.GetClient(stage).UserDelete(context.Background(), userName)
	if err != nil {
		return err
	}
	return err
}
