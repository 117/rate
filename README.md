# rate

```go
limiter := rate.NewLimiter(3, time.Second) // executions will be throttled to 3 per second

for {
  limiter.Sleep(func() {
      fmt.Println("only printed 3 times per second")
  })
}
```
