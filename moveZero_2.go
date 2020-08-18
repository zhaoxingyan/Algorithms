package main

import "fmt"

func main() {
	nums := []int{3, 0, 5, 0, 0, 8}
	fmt.Println(moveZeros(nums))
}

func moveZeros(nums []int) []int {
	len := len(nums)
	if len <= 1 {
		return nums
	}
	j := 0
	for i := 0; i < len; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < len; j++ {
		nums[j] = 0
	}
	return nums
}
