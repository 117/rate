package throttle

import (
	"fmt"
	"time"
)

var cooldowns map[string]time.Time

// Func -
func Func(key string, duration time.Duration, do func()) error {
	if _, has := cooldowns[key]; !has {
		cooldowns[key] = time.Now()
	}

	if time.Since(cooldowns[key]) > duration {
		cooldowns[key] = time.Now()
		do()
		return nil
	}

	return fmt.Errorf("key %s has been throttled", key)
}

// Key -
func Key(key string, duration time.Duration) error {
	if _, has := cooldowns[key]; !has {
		cooldowns[key] = time.Now()
	}

	if time.Since(cooldowns[key]) > duration {
		cooldowns[key] = time.Now()
		return nil
	}

	return fmt.Errorf("key %s has been throttled", key)
}
