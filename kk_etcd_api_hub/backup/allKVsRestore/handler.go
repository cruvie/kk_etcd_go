package allKVsRestore

func (x *api) Handler() (error, *AllKVsRestore_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return err, nil
	}
	return nil, &AllKVsRestore_Output{}
}
