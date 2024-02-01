## Слияние n каналов в один

На вход посутпает n каналов, необходимо слить их все в один результрующий канал и вывести его. 

## Реализация
```go
func MergeChan(chs ...<-chan int) <-chan int{
    var wg sync.WaitGroup
    res := make(chan int)
    for _, i := range chs {
        i := i
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := range i {
                res <- j
            }
        }()
        go func() {
            wg.Wait()
            close(res)
        }()
    }
    return res
}
```

## Пример
