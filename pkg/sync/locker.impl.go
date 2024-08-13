package sync

type fakeLocker struct{}

func (fl *fakeLocker) Lock() {
	// Fake Lock
}

func (fl *fakeLocker) Unlock() {
	// Fake Unlock
}

type fakeRWLocker struct{}

func (fl *fakeRWLocker) Lock() {
	// Fake Lock
}

func (fl *fakeRWLocker) Unlock() {
	// Fake Unlock
}

func (fl *fakeRWLocker) RLock() {
	// Fake RLock
}

func (fl *fakeRWLocker) RUnlock() {
	// Fake RUnlock
}

func (fl *fakeRWLocker) TryRLock() bool {
	return true
}

func (fl *fakeRWLocker) TryLock() bool {
	return true
}
