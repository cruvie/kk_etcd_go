package allKVsBackup

func (x *api) Handler() (*AllKVsBackup_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	response, err := x.service()
	if err != nil {
		return nil, err
	}
	return response, err
}
