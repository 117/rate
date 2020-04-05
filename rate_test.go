package rate

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLimiter(t *testing.T) {
	limiter := NewLimiter(3, time.Second)

	for {
		limiter.SleepIfRateExceeded(func() {
			fmt.Println(time.Now(), "should only be 3 per second")
		})
	}
}
