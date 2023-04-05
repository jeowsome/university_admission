package main

import (
	"fmt"
	"sort"
)

func main() {
	// DO NOT modify the code block below, it reads 4 int numbers from the stdin:
	var num1, num2, num3, num4 int
	fmt.Scanln(&num1, &num2, &num3, &num4)
	nums := []int{num1, num2, num3, num4}

	// Use the sort.SliceStable() function to sort the `nums` slice
	// by the remainder of dividing each number by 5 below:
	sort.SliceStable(nums, func(i, j int) bool {
		if nums[i] > nums[j] {
			return (nums[i] % 5) > (nums[j] % 5)
		}
		return (nums[i] % 5) < (nums[j] % 5)
	})

	fmt.Println(nums) // output the sorted slice here!
}
