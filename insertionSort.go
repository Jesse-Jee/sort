package sort

import "fmt"

// 稳定排序
// 时间复杂度O(N^2)
// 从前往后构建有序序列
// 对未排序数据，从已排序数据中从后向前扫描，找到相应位置插入。

func insertionSort(nums []int) {
	l := len(nums)
	var pre, cur int
	for i := 1; i < l; i++ {
		pre = i - 1
		cur = nums[i]
		for pre >= 0 && nums[pre] > cur {
			nums[pre+1] = nums[pre]
			pre--
		}
		nums[pre+1] = cur
	}

}

func main() {
	nums := []int{9, 8, 6, 7, 4, 5, 1, 3, 2}
	insertionSort(nums)
	fmt.Println(nums)
}
