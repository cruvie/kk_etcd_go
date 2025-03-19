package snapshotRestore

func (x *api) Handler() (error, *SnapshotRestore_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	cmdStr, err := x.service()
	if err != nil {
		return err, nil
	}
	return err, &SnapshotRestore_Output{CmdStr: cmdStr}
}
