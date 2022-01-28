package sync

import (
    "testing"
    "sync"
)

func TestCounter(t *testing.T) {
    t.Run("test increment behavior", func(t *testing.T) {
        // counter := Counter{};
        counter := NewCounter()
        counter.Inc()
        counter.Inc()
        counter.Inc()

        assertCounter(t, counter, 3)
        // if counter.Value() != 3 {
        //     t.Errorf("expected %d but got %d", 3, counter.Value())
        // }
    })

    t.Run("it runs safely concurrently", func(t *testing.T) {
        wantedCount := 1000
        // counter := Counter{}
        counter := NewCounter()
    
        var wg sync.WaitGroup
        wg.Add(wantedCount)
    
        for i := 0; i < wantedCount; i++ {
            go func() {
                counter.Inc()
                wg.Done()
            }()
        }
        wg.Wait()
    
        assertCounter(t, counter, wantedCount)
    })
}

func assertCounter(t testing.TB, got *Counter, want int)  {
    t.Helper()
    if got.Value() != want {
        t.Errorf("expected %d but got %d", want, got.Value())
    }
}

