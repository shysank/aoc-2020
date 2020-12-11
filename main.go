package main

import (
	"fmt"
)

func main() {

	var arr = []int{12, 13, 15, 16,17}



	w := ways(arr, len(arr)-1)

	fmt.Println(w)
}

func ways(arr []int, index int) int {
	if index == 0 {
		return 1
	}

	curr := arr[index]
	prevRoutes := ways(arr, index-1)
	var prev1, prev2 int
	if index-2 >= 0 && curr-arr[index-2] <= 3 {
		prev1 = ways(arr, index-2)
	}

	if index-3 >= 0 && curr-arr[index-3] <= 3 {
		prev2 = ways(arr, index-3)
	}

	return prevRoutes + prev1 + prev2

}
