package throttle

import (
	"fmt"
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	throttle := &Throttle{
		Limit:    3,
		Duration: time.Second,
	}

	for {
		throttle.Sleep(func() {
			fmt.Println(time.Now(), "should only be 3 per second")
		})
	}
}
