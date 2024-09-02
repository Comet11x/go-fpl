package events

import (
	"fmt"
	"testing"
)

func emit(ee EventEmitter) {
	ee.AsyncEmit(CreateEventWithoutPayload("TEST"))
}

func eventHandle(e Event) {
	fmt.Println("event: ", e.Name())
}

func TestCreateEE(t *testing.T) {
	ee := CreateEventEmitter()

	ee.On("TEST", CreateEventListener(func(e Event) {
		eventHandle(e)
		ee.RemoveAllEventListeners("TEST")
	}))

	go emit(ee)
	Awaiter(ee).Unwrap().Await()
}
