package events

func CreateAsyncModeEventPropagation() ModeEventPropagation {
	return &modeEventPropagation{m: _ASYNC_MODE}
}

func CreateSyncModeEventPropagation() ModeEventPropagation {
	return &modeEventPropagation{m: _SYNC_MODE}
}
