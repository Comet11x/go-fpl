package events

import "testing"

func TestCreateAsyncMode(t *testing.T) {
	m := CreateAsyncModeEventPropagation()

	if !m.IsAsync() || m.IsSync() {
		t.Fatal("It must be an asynchronous mode")
	}
}

func TestCreateSyncMode(t *testing.T) {
	m := CreateSyncModeEventPropagation()

	if !m.IsSync() || m.IsAsync() {
		t.Fatal("It must be a synchronous mode")
	}
}

func TestSyncModeIsEqualOtherSyncMode(t *testing.T) {
	m := CreateSyncModeEventPropagation()

	if !m.IsEqual(CreateSyncModeEventPropagation()) {
		t.Fatal("It must be equal")
	}
}

func TestSyncModeIsNotEqualAsyncMode(t *testing.T) {
	m := CreateSyncModeEventPropagation()

	if m.IsEqual(CreateAsyncModeEventPropagation()) {
		t.Fatal("It must be not equal")
	}
}

func TestAsyncModeIsEqualOtherAsyncMode(t *testing.T) {
	m := CreateAsyncModeEventPropagation()

	if !m.IsEqual(CreateAsyncModeEventPropagation()) {
		t.Fatal("It must be equal")
	}
}

func TestAsyncModeIsNotEqualSyncMode(t *testing.T) {
	m := CreateAsyncModeEventPropagation()

	if m.IsEqual(CreateSyncModeEventPropagation()) {
		t.Fatal("It must be not equal")
	}
}
