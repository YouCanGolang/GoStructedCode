package main

import "fmt"

// merge 函数用于合并两个有序数组
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// 比较左右两部分的元素，将较小的加入结果数组
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 将剩余的元素加入结果数组
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// mergeSort 实现归并排序
func mergeSort(arr []int) []int {
	// 如果数组长度小于2，直接返回
	if len(arr) < 2 {
		return arr
	}

	// 分解数组
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// 合并两个有序子数组
	return merge(left, right)
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	sortedArr := mergeSort(arr)
	fmt.Println("最终排序结果: ", sortedArr)
}
