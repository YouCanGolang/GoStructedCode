package main

import (
	"fmt"
)

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1[0])

	arrF()
}

func arrF() {
	var (
		arr1, arr2 [5]int
	)

	for i := 0; i < 5; i++ {
		arr1[i] = i
	}

	for i, v := range arr1 {
		arr2[5-i-1] = v
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
}
