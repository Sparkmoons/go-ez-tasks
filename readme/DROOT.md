## Нахождение цифрового корня числа

На вход подается число. Нужно найти его цифровой корень, то есть итеративно складывать цифры числа до тех пор, пока не останется только одна цифра. 

## Реализация 

```go
func cor(n int) int {
	for n > 9 {
		s := 0
		for n > 0 {
			s += n % 10
			n = n / 10
		}
		n = s
	}
	return n
}
```

## Пример

Пусть n = 65537. На выводе получим: 8
