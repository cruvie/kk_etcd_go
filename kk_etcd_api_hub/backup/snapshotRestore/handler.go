package snapshotRestore

func (x *api) Handler() (*SnapshotRestore_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	cmdStr, err := x.service()
	if err != nil {
		return nil, err
	}
	return &SnapshotRestore_Output{CmdStr: cmdStr}, err
}
