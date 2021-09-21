package knight

func (a *Api) StreamIncomingEvents() (*EventChan, error) {
	ec := &EventChan{
		channel: make(chan *Event),
		closed:  false,
	}
	r := new(Event)
	err := a.getEvent("/api/stream/event", nil, r, ec)
	if err != nil {
		return nil, err
	}
	return ec, nil
}
