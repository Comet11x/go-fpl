package events

type ModeEventPropagation interface {

	// Returns true if the mode is async
	IsAsync() bool

	// Returns true if the mode is sync
	IsSync() bool

	// Compares the mode and an other and returns true if they both sync or async
	IsEqual(other ModeEventPropagation) bool
}
