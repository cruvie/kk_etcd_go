package kVPut

func (x *api) Handler() (error, *KVPut_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return err, &KVPut_Output{}
}
