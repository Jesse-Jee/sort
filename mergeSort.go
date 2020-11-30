package sort

import "fmt"

// 稳定排序
// 时间复杂度O(NlogN)
// 分治法
//	将已有序的子序列合并，得到完全有序的序列。
//
//

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	l := right-left+1
	tmp := make([]int,l)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}

	for p := 0; p < len(tmp); p++ {
		nums[left+p] = tmp[p]
	}
}

func main() {
	nums := []int{9, 8, 6, 7, 4, 5, 1, 3, 2}
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
