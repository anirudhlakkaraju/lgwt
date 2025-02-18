package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment three times and leave it there", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter.Value(), 3)
	})
}

func assertCounter(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", want, got)
	}
}
