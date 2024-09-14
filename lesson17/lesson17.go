package main

import "fmt"

func bubbleSort(arr []int) {
	length := len(arr)
	num := 0

	// 轮次控制
	for i := 0; i < length-1; i++ {
		fmt.Printf("\n第 %d 轮\n", i+1)

		// 比较次数控制
		for j := 0; j < length-i-1; j++ {
			num++
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Printf("第 %d 次比较，结果: %v\n", num, arr)
		}
	}
}

func bubbleSortOpm(arr []int) {
	length := len(arr)
	num := 0

	// 轮次控制
	for i := 0; i < length-1; i++ {
		fmt.Printf("\n第 %d 轮\n", i+1)

		// 此次是否进行过交换
		swapped := false

		// 比较次数控制
		for j := 0; j < length-i-1; j++ {
			num++
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Printf("第 %d 次比较，结果: %v\n", num, arr)
		}

		// 如果没有发生交换，说明数组已经有序，提前退出
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	bubbleSort(arr)
	fmt.Println("排序结果: ", arr)
}
