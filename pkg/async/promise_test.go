package async

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestPromise(t *testing.T) {
	done := atomic.Bool{}
	promise := Async[string](func(resolve func(string), reject func(any)) {
		time.Sleep(time.Second * 2)
		resolve("Hola")
	})

	promise.Then(func(s string) {
		done.Store(true)
		t.Logf("This handler receives a message: '%s'", s)
	})

	promise.Await()

	time.Sleep(time.Second * 3)
	promise.Then(func(s string) {
		t.Logf("Another handler receives a message: '%s'", s)
	})
	if !done.Load() {
		t.Fatal("It must be true")
	}
}
