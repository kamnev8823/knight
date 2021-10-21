package knight

type streamInterface interface {
	close()
	write(interface{}) bool
}

type EventChan struct {
	channel chan *Event
	signal  chan bool
}

func (ec *EventChan) close() {
	ec.signal <- true
	close(ec.channel)
}

func (ec *EventChan) write(c interface{}) bool {
	select {
	case <-ec.signal:
		return false
	default:
		switch c.(type) {
		case *Event:
			ec.channel <- c.(*Event)
			return true
		}
	}

	return false
}

type TVChan struct {
	channel chan *TVStream
	signal  chan bool
}

func (tvc *TVChan) close() {
	tvc.signal <- true
	close(tvc.channel)
}

func (tvc *TVChan) write(c interface{}) bool {
	select {
	case <-tvc.signal:
		return false
	default:
		switch c.(type) {
		case *TVStream:
			tvc.channel <- c.(*TVStream)
			return true
		}
	}
	return false
}
