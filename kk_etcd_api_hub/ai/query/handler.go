package query

func (x *api) Handler() (error, *Query_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	answer, err := x.service()
	if err != nil {
		return err, nil
	}
	return nil, &Query_Output{
		Answer: answer,
	}
}
