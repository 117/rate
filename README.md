# rate

```go
<<<<<<< HEAD
func main() {
  limiter := rate.NewLimiter(3, time.Second)

  for {
    limiter.Sleep(func() {
      fmt.Println("only printed 3 times per second")
    })
  }
=======
limiter := rate.NewLimiter(3, time.Second) // executions will be throttled to 3 per second

for {
  limiter.Sleep(func() {
      fmt.Println("only printed 3 times per second")
  })
>>>>>>> 4259e27784d7b49e90fb63dfd423a10ac095afc9
}
```
