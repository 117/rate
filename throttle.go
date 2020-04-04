package throttle

import (
	"errors"
	"time"
)

// Throttle is a wrapper for limiting the number of executions for a specified rolling duration.
type Throttle struct {
	Duration time.Duration
	Limit    int

	stamp  time.Time
	called int
}

// Error returns an error if execution has been throttled.
func (t *Throttle) Error(exec func()) error {
	if t.stamp.IsZero() || t.Duration-time.Since(t.stamp) < 0 {
		t.stamp = time.Now()
		t.called = 0
	}

	if t.called >= t.Limit {
		return errors.New("throttle was triggered")
	}

	exec()
	t.called++
	return nil
}

// Sleep blocks until the throttle is over.
func (t *Throttle) Sleep(exec func()) {
	if t.Error(exec) != nil {
		time.Sleep(t.Duration - time.Since(t.stamp))
	}
}
