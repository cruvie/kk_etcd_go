package snapshot

func (x *api) Handler() (error, *Snapshot_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err, response := x.service()
	if err != nil {
		return err, nil
	}
	return err, response
}
