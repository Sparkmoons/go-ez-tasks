## Конвейер чисел

На примере до 10. В первый канал принимаются 10 чисел, затем второй канал возвращает квадрат каждого числа.

## Реализация

```go
func main(){
  var wg sync.WaitGroup
  inp := make(chan int)
  out := make(chan int)

  go func() {
    for i := 0; i < 10; i++ {
      inp <- i
    }
    close(inp)
  }()

  go func() {
    for j := range inp {
      out <- j * j
    }
    close(out)
  }()

  wg.Add(1)
  go func() {
    defer wg.Done()
    for x := range out {
      fmt.Println(x)
  }     
  }()
  wg.Wait()
}
```

## Пример

