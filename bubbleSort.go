package sort

import "fmt"

// 稳定的排序
// 时间复杂度O(N^2)
// 相邻元素比较，如果顺序错误就交换

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

func main() {
	nums := []int{9, 8, 6, 7, 4, 5, 1, 3, 2}
	bubbleSort(nums)
	fmt.Println(nums)
}
