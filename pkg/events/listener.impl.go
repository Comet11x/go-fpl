package events

type eventListener struct {
	t  int64
	fn func(Event)
}

func (l eventListener) Id() int64 {
	return l.t
}

func (l *eventListener) Call(e Event) {
	l.fn(e)
}
