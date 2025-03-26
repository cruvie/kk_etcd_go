package snapshotInfo

func (x *api) Handler() (*SnapshotInfo_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	info, err := x.service()
	if err != nil {
		return nil, err
	}
	return &SnapshotInfo_Output{Info: info}, err
}
