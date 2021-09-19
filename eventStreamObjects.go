package knight

type streamEventInterface interface {
	Close()
	isClosed() bool
	write(interface{}) bool
}

type EventChan struct {
	channel chan *Event
	closed  bool
}

func (ec *EventChan) GetChan() <-chan *Event {
	return ec.channel
}

func (ec *EventChan) Close() {
	if !ec.closed {
		ec.closed = true
		close(ec.channel)
	}
}

func (ec *EventChan) write(c interface{}) bool {
	if !ec.closed {
		switch c.(type) {
		case *Event:
			ec.channel <- c.(*Event)
			return true
		}
	}
	return false
}

func (ec *EventChan) isClosed() bool {
	return ec.closed
}
