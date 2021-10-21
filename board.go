package knight

func (a *Api) StreamIncomingEvents() (<-chan *Event, error) {
	ec := &EventChan{
		channel: make(chan *Event),
		signal:  make(chan bool, 1),
	}
	r := new(Event)
	err := a.getEvent("/api/stream/event", nil, r, ec)
	if err != nil {
		return nil, err
	}
	return ec.channel, nil
}
