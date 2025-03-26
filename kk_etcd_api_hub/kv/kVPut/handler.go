package kVPut

func (x *api) Handler() (*KVPut_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return &KVPut_Output{}, err
}
