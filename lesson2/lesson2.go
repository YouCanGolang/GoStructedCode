package main

import "fmt"

func main() {
	// 声明长度为 5 的整数数组
	var (
		arr1 [5]int
	)
	fmt.Println(len(arr1))

	// 通过字面量初始化数组
	arr2 := [5]int{1, 2, 3, 4, 5}

	// 自动推断数组长度
	arr3 := [...]int{10, 20, 30}
	fmt.Println(len(arr3))

	// 访问数组中的元素
	fmt.Println(arr2[2]) // 输出：3

	// 修改数组中的元素
	arr2[2] = 10
	fmt.Println(arr2[2]) // 输出：10

	// 通过索引遍历数组
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// 使用 range 遍历数组
	for i, v := range arr2 {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}

	fmt.Println("最大值为：", findMax(arr2))

	fmt.Println("反转后的数组：", reverseArray(arr2))
}

func findMax(arr [5]int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func reverseArray(arr [5]int) [5]int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}
