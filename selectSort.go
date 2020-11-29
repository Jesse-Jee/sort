package sort

import "fmt"

// 比较类排序
// 不稳定
// 时间复杂度O(N^2)
// 每次选择最小值，放到待排序数组的起始位置。

func SelectSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		var min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

func main() {
	nums := []int{9, 8, 6, 7, 4, 5, 1, 3, 2}
	sorted := SelectSort(nums)
	fmt.Println(sorted)
}
