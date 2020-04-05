# rate

```go
func main() {
  limiter := rate.NewLimiter(3, time.Second)

  for {
    limiter.SleepIfRateExceeded(func() {
      fmt.Println("only printed 3 times per second")
    })
  }
}
```
