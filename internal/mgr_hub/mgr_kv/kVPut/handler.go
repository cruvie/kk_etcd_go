package kVPut

func (x *api) handler() (error, *KVPut_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return err, &KVPut_Output{}
}
