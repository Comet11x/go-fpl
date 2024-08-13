package events

const (
	_ASYNC_MODE = 0
	_SYNC_MODE  = 1
)

type modeEventPropagation struct {
	m uint8
}

func (m *modeEventPropagation) IsAsync() bool {
	return m.m == _ASYNC_MODE
}

func (m *modeEventPropagation) IsSync() bool {
	return m.m == _SYNC_MODE
}

func (m *modeEventPropagation) IsEqual(other ModeEventPropagation) bool {
	return m.IsAsync() && other.IsAsync() || m.IsSync() && other.IsSync()
}
