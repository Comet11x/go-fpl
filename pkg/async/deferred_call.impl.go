package async

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/comet11x/go-fpl/pkg/core"
)

const (
	_PENDING  = 0
	_DONE     = 1
	_CANCELED = 2
)

type deferredCall[T any] struct {
	timeout    time.Duration
	try        core.Try[T]
	out        core.Option[core.Try[T]]
	mutex      sync.Mutex
	isCanceled atomic.Bool
	status     uint8
	awaiter    chan struct{}
}

func (i *deferredCall[T]) call() {
	time.Sleep(i.timeout)
	i.mutex.Lock()
	if i.status == _PENDING {
		i.status = _DONE
		i.try.IsSuccess()
		i.out = core.Some(i.try)
		if i.awaiter != nil {
			i.awaiter <- struct{}{}
		}
	}
	i.mutex.Unlock()
}

func (i *deferredCall[T]) Cancel() {
	i.mutex.Lock()
	if i.status == _PENDING {
		i.isCanceled.Store(true)
		i.status = _CANCELED
		i.out = core.None[core.Try[T]]()
		if i.awaiter != nil {
			i.awaiter <- struct{}{}
		}
	}
	i.mutex.Unlock()
}

func (i *deferredCall[T]) IsCanceled() bool {
	return i.isCanceled.Load()
}

func (i *deferredCall[T]) Await() core.Either[core.Option[core.Try[T]], any] {
	i.mutex.Lock()
	if i.out != nil {
		return core.Left[core.Option[core.Try[T]], any](i.out)
	} else {
		i.awaiter = make(chan struct{})
	}
	i.mutex.Unlock()

	<-i.awaiter
	return core.Left[core.Option[core.Try[T]], any](i.out)
}
