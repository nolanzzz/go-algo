package QuickSort

import (
	"fmt"
	"testing"
)

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	pivot := nums[left]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func TestQuickSort(t *testing.T) {
	nums := []int{3, 1, 6, 5, 8, 10, 15, 9, 10, 100, 6, 99, -3}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
