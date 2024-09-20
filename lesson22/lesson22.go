package main

import (
	"fmt"
)

// linearSearch 实现线性查找
func linearSearch(arr []int, target int) int {
	// 遍历数组
	for i, val := range arr {
		// 如果找到目标值，返回其索引
		if val == target {
			return i
		}
	}
	// 如果找不到，返回 -1
	return -1
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	target := 4

	// 执行线性查找
	result := linearSearch(arr, target)

	if result != -1 {
		fmt.Printf("元素 %d 找到，索引为: %d\n", target, result)
	} else {
		fmt.Println("元素未找到")
	}
}
