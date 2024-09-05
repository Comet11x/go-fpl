package async

import (
	"sync/atomic"
	"testing"
	"time"
)

func fork[T any](p Promise[T], logf func(format string, args ...any)) {
	either := p.Await()
	either.IfLeft(func(value T) {
		logf("2nd goroutine says %s", value)
	})
}

func TestPromise(t *testing.T) {

	// if this test is finished successfully it will have true
	done := atomic.Bool{}

	// It creates a promise
	promise := Async(func(resolve func(string), reject func(any)) {
		// This function will be called asynchronously
		msg := "Hola"
		t.Logf("Async function says '%s'", msg)
		time.Sleep(time.Second * 2)
		resolve(msg)
	})

	// It adds 1st event handler
	promise.Then(func(s string) {
		done.Store(true)
		t.Logf("1st handler has received a message '%s'", s)
	})

	// It sends promise to 2nd goroutine
	go fork(promise, t.Logf)

	t.Logf("1st goroutine is going to sleep 2 seconds")
	time.Sleep(time.Second * 3)
	t.Logf("1st goroutine has woken up")

	// Waiting for results from the async function
	promise.Await().
		IfLeft(func(value string) {
			// The result has received successfully
			t.Logf("2nd goroutine says '%s'", value)
		}).
		IfRight(func(value any) {
			// An abnormal result
			t.Fatal("Something went wrong! It has received an abnormal value: ", value)
		})

	time.Sleep(time.Second * 3)

	// It adds another event handler
	promise.Then(func(s string) {
		t.Logf("2nd handler has received a message '%s'", s)
	})

	if !done.Load() {
		t.Fatal("It must be true")
	}
}
