package main

import "fmt"

func insertionSort(arr []int) {
	// 从第二个元素开始遍历，认为第一个元素是已排序的
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// 在已排序部分中寻找插入位置
		for j >= 0 && arr[j] > key {
			// 将比key大的元素右移
			arr[j+1] = arr[j]
			j--
		}
		// 插入key到正确位置
		arr[j+1] = key

		fmt.Printf("第 %d 轮插入结果: %v\n", i, arr)
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	insertionSort(arr)
	fmt.Println("排序结果: ", arr)
}
