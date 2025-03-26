package kVDel

func (x *api) Handler() (*KVDel_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return &KVDel_Output{}, err

}
