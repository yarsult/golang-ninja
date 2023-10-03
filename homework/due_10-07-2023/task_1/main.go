package main

import "fmt"

func fibonacci() func(int) int {
	nums := []int{0, 1, 1}
	return func(n int) int {
		if n < 3 {
			return nums[n]
		}
		nums = nums[:n]
		nums = append(nums, nums[n-1]+nums[n-2])
		return nums[n]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
