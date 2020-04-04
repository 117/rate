// @module github.com/117/rate@v0.0.3

package rate

import (
	"errors"
	"time"
)

// Limiter is a wrapper for limiting the number of executions for a specified rolling duration.
type Limiter struct {
	Duration time.Duration
	Limit    int
	stamp    time.Time
	called   int
}

// Error returns an error if the call limit for the duration has been reached.
func (l *Limiter) Error(exec func()) error {
	if l.stamp.IsZero() || l.Duration-time.Since(l.stamp) < 0 {
		l.stamp = time.Now()
		l.called = 0
	}

	if l.called >= l.Limit {
		return errors.New("call limit reached")
	}

	exec()
	l.called++
	return nil
}

// Sleep blocks until the throttle is over.
func (l *Limiter) Sleep(exec func()) {
	if l.Error(exec) != nil {
		time.Sleep(l.Duration - time.Since(l.stamp))
	}
}

// NewLimiter creates a new limiter with the provided call limit and duration.
func NewLimiter(limit int, duration time.Duration) *Limiter {
	return &Limiter{
		Limit:    limit,
		Duration: duration,
	}
}
