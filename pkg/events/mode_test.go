package events

import "testing"

func TestCreateAsyncMode(t *testing.T) {
	m := AsyncModeEventPropagation()

	if !m.IsAsync() || m.IsSync() {
		t.Fatal("It must be an asynchronous mode")
	}
}

func TestCreateSyncMode(t *testing.T) {
	m := SyncModeEventPropagation()

	if !m.IsSync() || m.IsAsync() {
		t.Fatal("It must be a synchronous mode")
	}
}

func TestSyncModeIsEqualOtherSyncMode(t *testing.T) {
	m := SyncModeEventPropagation()

	if !m.IsEqual(SyncModeEventPropagation()) {
		t.Fatal("It must be equal")
	}
}

func TestSyncModeIsNotEqualAsyncMode(t *testing.T) {
	m := SyncModeEventPropagation()

	if m.IsEqual(AsyncModeEventPropagation()) {
		t.Fatal("It must be not equal")
	}
}

func TestAsyncModeIsEqualOtherAsyncMode(t *testing.T) {
	m := AsyncModeEventPropagation()

	if !m.IsEqual(AsyncModeEventPropagation()) {
		t.Fatal("It must be equal")
	}
}

func TestAsyncModeIsNotEqualSyncMode(t *testing.T) {
	m := AsyncModeEventPropagation()

	if m.IsEqual(SyncModeEventPropagation()) {
		t.Fatal("It must be not equal")
	}
}
