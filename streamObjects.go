package knight

type streamInterface interface {
	close()
	write(interface{}) bool
}

type EventChan struct {
	channel chan *Event
}

func (ec *EventChan) close() {
	close(ec.channel)
}

func (ec *EventChan) write(c interface{}) bool {
	switch c.(type) {
	case *Event:
		ec.channel <- c.(*Event)
		return true
	}
	return false
}

type TVChan struct {
	channel chan *TVStream
}

func (tvc *TVChan) close() {
	close(tvc.channel)
}

func (tvc *TVChan) write(c interface{}) bool {
	switch c.(type) {
	case *TVStream:
		tvc.channel <- c.(*TVStream)
		return true
	}
	return false
}
