package syncs

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times, leave it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely in concurent mode", func(t *testing.T) {
		wantedCounte := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCounte)

		for i := 0; i < wantedCounte; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, counter, wantedCounte)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
