package main

import (
	"fmt"
)

// partition 函数实现数组的划分
func partition(arr []int, low, high int) int {
	pivot := arr[high] // 选择最后一个元素作为基准
	i := low - 1       // i 代表已处理的元素区间

	// 遍历数组，将小于 pivot 的元素移到前面
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 将基准元素放到正确位置
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // 返回基准元素的位置
}

// quickSort 实现快速排序
func quickSort(arr []int, low, high int) {
	if low < high {
		// 划分数组，并获取基准元素的位置
		pi := partition(arr, low, high)

		// 递归排序基准左侧和右侧的子数组
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("最终排序结果: ", arr)
}
