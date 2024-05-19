## Сложение бинарных чисел

Задача с литкода. Подаются два бинарных числа, нужна написать функцию, которая осуществляет их сложение.

## Реализация

```go
func addBinary(a string, b string) string {
    res := ""
    i, j := len(a)-1, len(b)-1
    last := 0

    for i >= 0 || j >= 0 || last > 0 {
        sum := 0
        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }
        sum += last
        last = sum / 2
        res = string(sum % 2 + '0') + res
    }
    return res
}
```

## Пример
Пусть a = "11", b = "1"

На выходе будет: "100"
