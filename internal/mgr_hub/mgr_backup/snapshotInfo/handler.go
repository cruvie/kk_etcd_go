package snapshotInfo

func (x *api) handler() (error, *SnapshotInfo_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	info, err := x.service()
	if err != nil {
		return err, nil
	}
	return err, &SnapshotInfo_Output{Info: info}
}
