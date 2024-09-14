package main

import "fmt"

func selectionSort(arr []int) {
	length := len(arr)
	num := 0

	// 轮次控制
	for i := 0; i < length-1; i++ {
		fmt.Printf("\n第 %d 轮\n", i+1)

		// 假设当前未排序部分的第一个元素是最小值
		minIndex := i

		// 找到未排序部分的最小元素
		for j := i + 1; j < length; j++ {
			num++
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			fmt.Printf("第 %d 次比较\n", num)
		}

		fmt.Printf("此轮已排序序列：%v, ", arr[0:i])
		fmt.Printf("此轮未排序序列：%v, ", arr[i:length])

		// 将未排序部分的最小值与当前元素交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		fmt.Printf("此轮比较结果: %v\n", arr)
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	selectionSort(arr)
	fmt.Println("\n排序结果: ", arr)
}
