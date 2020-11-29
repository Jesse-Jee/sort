package sort

import "fmt"

// 不稳定的排序
// 时间复杂度O(NlogN)
// 分治思想
// 数组取pivot，小于pivot的元素放左边，大于它的放右边。
// 然后依次对左边和右边的子数组进行快排，以达到整个序列有序。

func quickSort(nums []int, begin, end int) {
	if begin > end {
		return
	}
	pivot := begin
	i := pivot + 1

	for j := begin; j <= end; j++ {
		if nums[j] < nums[pivot] {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}

	nums[i-1], nums[pivot] = nums[pivot], nums[i-1]

	quickSort(nums, begin, i-2)
	quickSort(nums, i, end)

}

func main() {
	nums := []int{9, 8, 6, 7, 4, 5, 1, 3, 2}
	quickSort(nums,0,len(nums)-1)
	fmt.Println(nums)
}
