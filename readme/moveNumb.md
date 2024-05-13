## Смещение чисел в слайсе

Дается числовой слайс, необходимо переместить определенные числа, например 0, в конец слайса, при этом сохранив порядок других чисел

## Реализация

```go
func moveZeroes(nums []int) []int {
	nonZeroIdx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIdx], nums[i] = nums[i], nums[nonZeroIdx]
			nonZeroIdx++
		}
	}
	return nums
}

```