// 代码生成时间: 2025-10-30 23:09:19
package main

import (
    "buffalo"
    "buffalo/worker"
    "github.com/markbates/inflect"
    "log"
    "net/http"
    "os"
    "time"
    "strings"
)

// StressTester represents the structure for stress testing
type StressTester struct {
    URL      string
    Workers  int
    Duration time.Duration
}

// NewStressTester creates a new StressTester instance
func NewStressTester(url string, workers int, duration time.Duration) *StressTester {
    return &StressTester{
        URL:      url,
        Workers:  workers,
        Duration: duration,
    }
}

// Start starts the stress test
func (s *StressTester) Start() error {
    log.Printf("Starting stress test on %s with %d workers for %s", s.URL, s.Workers, s.Duration)

    // Create a worker pool
    pool := worker.NewPool(s.Workers)
    defer pool.Close()

    start := time.Now()
    for time.Since(start) < s.Duration {
        // Enqueue a new worker task
        _, err := pool.Enqueue(func(workers *worker.Pool) {
            _, err := http.Get(s.URL)
            if err != nil {
                log.Printf("Error during request: %v", err)
                return
            }
        })
        if err != nil {
            log.Printf("Error enqueuing worker: %v", err)
        return err
    }

        // Sleep for a short duration to simulate realistic request rate
        time.Sleep(1 * time.Second)
    }

    return nil
}

func main() {
    if len(os.Args) < 4 {
        log.Fatalf("Usage: %s <URL> <Workers> <Duration (s)>

