package events

import (
	"fmt"
	"testing"
	"time"
)

const (
	HELLO_EVENT_NAME = "hello"
	BYE_EVENT_NAME   = "exit"
)

func async(ee EventEmitter) {
	time.Sleep(time.Second * 2)
	ee.Emit(ee.CreateEvent(
		HELLO_EVENT_NAME,
		EventOptions("", false).
			SetPayload("Hello World!").
			SetSenderName("TEST"),
	))
	// time.Sleep(time.Second * 2)
	ee.Emit(ee.CreateEvent(
		BYE_EVENT_NAME,
		EventOptions("", false).SetPayload("Goodbye World").SetSenderName("EXIT"),
	))
	time.Sleep(time.Second * 2)
}

func TestEE(t *testing.T) {
	ee := New()
	var ls [2]EventListener
	ls[0] = ee.CreateListener(func(e Event) {
		fmt.Println(e.From(), " :> ", e.Payload().(string))
	})

	ls[1] = ee.CreateListener(func(e Event) {
		fmt.Println(e.From(), " :<> ", e.Payload().(string))
		time.Sleep(time.Second * 2)
		t.Log("Exit")
		for i, s := range []string{HELLO_EVENT_NAME, BYE_EVENT_NAME} {
			ee.Off(s, ls[i])
		}
	})

	ee.On(HELLO_EVENT_NAME, ls[0])
	ee.On(BYE_EVENT_NAME, ls[1])
	go async(ee)
	ee.Wait()
	fmt.Printf("EXIT FROM TEST")
}
