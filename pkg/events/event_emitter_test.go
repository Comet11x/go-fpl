package events

import (
	"fmt"
	"testing"

	"github.com/comet11x/go-fpl/pkg/types"
)

func emit(ee EventEmitter) {
	ee.AsyncEmit(EventWithoutPayload("TEST"))
}

func eventHandle(e Event) {
	fmt.Println("event: ", e.Name())
}

func TestCreateEE(t *testing.T) {
	ee := NewEventEmitter()

	// add an event listener
	ee.On("TEST", NewEventListener(func(e Event) {
		eventHandle(e)
		ee.RemoveAllEventListeners("TEST")
	}))

	go emit(ee)

	TryAwaiterFrom(ee).
		IfOk(func(value types.Awaiter[types.Void]) {
			t.Log("All listeners have unsubscribed")
		})
}
