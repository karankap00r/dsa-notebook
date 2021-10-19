package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Print("The majority element is: ", getMajorityElement(nums))
}

func getMajorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	result := math.MinInt
	cnt := 0
	for _, val := range nums {
		if val == result {
			cnt++
		} else {
			if cnt > 0 {
				cnt--
			} else {
				result = val
				cnt = 1
			}
		}
	}

	cnt = 0
	for _, val := range nums {
		if val == result {
			cnt++
		}
	}

	if cnt > len(nums)/2 {
		return result
	} else {
		return math.MinInt
	}
}
