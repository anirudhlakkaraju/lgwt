package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment three times and leave it there", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 7000
		counter := NewCounter()

		var wg sync.WaitGroup

		// Add count of dones
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()

				// Wait till goroutine is complete
				wg.Done()

			}()
		}
		// Wait until all goroutines are complete
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.value != want {
		t.Errorf("got %d, want %d", got.value, want)
	}
}
