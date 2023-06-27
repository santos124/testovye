
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界", ls([]int{1, 1, 0, 0, 1, 0, 1, 0, 0, 1}), 2)
}

//[1,1,0,0,1,0,1,0,0,1]

func ls(nums []int) int {
	maxInterval := 0
	currentInterval := 0
	score := 1
	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
			currentInterval++
		} else if (score == 1 && currentInterval > 0 && i < len(nums)-1 && nums[i+1]==1) {
			cnt++
			currentInterval++
			i++
			score = 0
		} else if currentInterval > maxInterval {
			maxInterval = currentInterval
			currentInterval = 0
			score = 1
		} else {
			currentInterval = 0
			score = 1
		}
	}
	if currentInterval > maxInterval {
		maxInterval = currentInterval
	}
	if cnt == len(nums) {
		return len(nums) - 1
	}
	return maxInterval
}
