package allKVsRestore

func (x *api) Handler() (*AllKVsRestore_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return nil, err
	}
	return &AllKVsRestore_Output{}, nil
}
