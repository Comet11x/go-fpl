package sync

import "sync"

type Locker interface {
	sync.Locker
}

type RWLocker interface {
	sync.Locker
	RLock()
	RUnlock()
	TryRLock() bool
	TryLock() bool
}
