package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)

	length := 0
	max_length := 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i + 1] == nums[i] {
			continue
		}
		if nums[i + 1] - nums[i] != 1 {
			if length > max_length {
				max_length = length
			}
			length = 0
		} else {
			length++
		}
	}
	if length > max_length {
		max_length = length
	}

	fmt.Println(max_length + 1)
}
