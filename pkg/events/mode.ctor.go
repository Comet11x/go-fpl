package events

func AsyncModeEventPropagation() ModeEventPropagation {
	return &modeEventPropagation{m: _ASYNC_MODE}
}

func SyncModeEventPropagation() ModeEventPropagation {
	return &modeEventPropagation{m: _SYNC_MODE}
}
