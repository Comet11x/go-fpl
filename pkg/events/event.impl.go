package events

import (
	"time"

	"github.com/comet11x/go-fpl/pkg/core"
)

type event struct {
	name       string
	cancelable bool
	canceled   bool
	t          time.Time
	from       core.Option[any]
	payload    core.Option[any]
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
	if e.cancelable {
		e.canceled = true
	}
}

func (e *event) From() core.Option[any] {
	return e.from
}

func (e *event) Time() time.Time {
	return e.t
}

func (e *event) Payload() core.Option[any] {
	return e.payload
}
