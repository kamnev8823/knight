package knight

type streamInterface interface {
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

type TVChan struct {
	channel chan *TVStream
	closed  bool
}

func (tvc *TVChan) GetChan() <-chan *TVStream {
	return tvc.channel
}

func (tvc *TVChan) Close() {
	if !tvc.closed {
		tvc.closed = true
		close(tvc.channel)
	}
}

func (tvc *TVChan) write(c interface{}) bool {
	if !tvc.closed {
		switch c.(type) {
		case *TVStream:
			tvc.channel <- c.(*TVStream)
			return true
		}
	}
	return false
}

func (tvc *TVChan) isClosed() bool {
	return tvc.closed
}
