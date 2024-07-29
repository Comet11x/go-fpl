package events

import (
	"time"
)

type event struct {
	cancelable bool
	canceled   bool
	stopped    bool
	time       time.Time
	from       any
	name       string
	payload    any
	async      bool
}

func (e *event) Name() string {
	return e.name
}

func (e *event) IsCancelable() bool {
	return e.cancelable
}

func (e *event) IsCanceled() bool {
	return e.canceled
}

func (e *event) Cancel() {
	if e.cancelable && !e.canceled {
		e.canceled = true
	}
}

func (e *event) From() any {
	return e.from
}

func (e *event) Time() time.Time {
	return e.time
}

func (e *event) Payload() any {
	return e.payload
}
