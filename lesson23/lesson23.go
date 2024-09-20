package main

import (
	"fmt"
)

// binarySearch 实现二分查找
func binarySearch(arr []int, target int) int {
	low := 0             // 查找范围的起始点
	high := len(arr) - 1 // 查找范围的结束点

	// 当查找范围不为空时
	for low <= high {
		mid := low + (high-low)/2 // 计算中间元素索引

		// 如果目标值与中间元素相等，查找成功
		if arr[mid] == target {
			return mid
		}

		// 如果目标值小于中间元素，缩小查找范围到左半部分
		if target < arr[mid] {
			high = mid - 1
		} else {
			// 如果目标值大于中间元素，缩小查找范围到右半部分
			low = mid + 1
		}
	}

	// 如果找不到，返回 -1
	return -1
}

// binarySearchRecursive 递归实现二分查找
func binarySearchRecursive(arr []int, low int, high int, target int) int {
	// 基线条件，查找范围为空
	if low > high {
		return -1
	}

	// 计算中间元素索引
	mid := low + (high-low)/2

	// 查找成功
	if arr[mid] == target {
		return mid
	}

	// 目标值小于中间元素，递归查找左半部分
	if target < arr[mid] {
		return binarySearchRecursive(arr, low, mid-1, target)
	}

	// 目标值大于中间元素，递归查找右半部分
	return binarySearchRecursive(arr, mid+1, high, target)
}

func main() {
	arr := []int{2, 3, 4, 5, 8}
	target := 4

	// 执行二分查找
	result := binarySearch(arr, target)

	if result != -1 {
		fmt.Printf("元素 %d 找到，索引为: %d\n", target, result)
	} else {
		fmt.Println("元素未找到")
	}
}
