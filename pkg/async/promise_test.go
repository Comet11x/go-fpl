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
		resolve("Hello")
	})

	promise.Then(func(s string) {
		done.Store(true)
		t.Log("A value from the promise: ", s)
	})

	promise.Await()

	// time.Sleep(time.Second * 3)
	if !done.Load() {
		t.Fatal("it must be true")
	}
}
