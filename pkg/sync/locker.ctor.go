package sync

import "sync"

func FakeRWLocker() RWLocker {
	return &fakeRWLocker{}
}

func FakeLocker() Locker {
	return &fakeLocker{}
}

func RealRWLocker() RWLocker {
	return &sync.RWMutex{}
}

func RealLocker() Locker {
	return &sync.Mutex{}
}
