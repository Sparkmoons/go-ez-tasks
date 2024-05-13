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