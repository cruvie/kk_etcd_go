package snapshot

func (x *api) Handler() (*Snapshot_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	response, err := x.service()
	if err != nil {
		return nil, err
	}
	return response, err
}
