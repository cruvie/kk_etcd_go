package query

func (x *api) Handler() (*Query_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	answer, err := x.service()
	if err != nil {
		return nil, err
	}
	return &Query_Output{
		Answer: answer,
	}, nil
}
