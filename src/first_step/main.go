package main

import (
	"fmt"
)

func FindLargest(nums []int) int {
	largest := nums[0]
	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}

	fmt.Println(largest)
	return largest

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	FindLargest(nums)
}
