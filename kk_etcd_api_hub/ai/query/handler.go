package query

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/ai/api_def"

func (x *api) Handler() (*api_def.Query_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	answer, err := x.service()
	if err != nil {
		return nil, err
	}
	return &api_def.Query_Output{
		Answer: answer,
	}, nil
}
