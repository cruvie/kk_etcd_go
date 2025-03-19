package kVDel

func (x *api) Handler() (error, *KVDel_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return err, &KVDel_Output{}

}
