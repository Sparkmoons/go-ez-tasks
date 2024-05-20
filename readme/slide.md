## Поиск подстроки

Необходимо найти в имеющейся строке самую длинную подстроку, содержащую только уникальные элементы. Алгоритм со сложность O(N).

## Реализация

```go
func longest(s string) float64 {

	left, right := 0, 0
	answer := 0.0
	set := make(map[byte]bool)

	for right < len(s) {
		c := s[right]
		if !set[c] {
			set[c] = true
			answer = math.Max(answer, float64(right-left+1))
			right++
		} else {
			for set[c] {
				delete(set, s[left])
				left++
			}
		}
	}
	return answer
}
```

## Пример

На вход подается строка: "abcdebed". На выводе получим 5.
