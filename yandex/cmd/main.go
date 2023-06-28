package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	
	fmt.Println("Hello, 世界", ls([]int{1,0,1,1,0,1,1}), 4)

	fmt.Println("Hello, 世界", ls([]int{1,0,1,1,1,1,1,1,0,1,1,1,1,1}), 11)

	fmt.Println("Hello, 世界", ls([]int{1, 1, 0, 0, 1, 0, 1, 0, 0, 1}), 2)
}

//[1,1,0,0,1,0,1,0,0,1]

func ls(nums []int) int {
	maxInterval := 0
	currentInterval := 0
	score := 1
	cnt := 0

	for i := range nums {
		if nums[i] ==1 {
			cnt++
		}
	}
	checkPoint := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currentInterval++
		} else if (score == 1 && currentInterval > 0 && i < len(nums)-1 && nums[i+1]==1) {
			currentInterval++
			checkPoint = i
			i++
			score = 0
		} else if currentInterval > maxInterval {
			maxInterval = currentInterval
			currentInterval = 0
			score = 1
			if checkPoint != -1 {
				i = checkPoint
				checkPoint = -1
			}
		} else {
			if checkPoint != -1 {
				i = checkPoint
				checkPoint = -1
			}
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

func asd() {
	wg := &sync.WaitGroup{}
	for i:= 0; i < 100; i++ {
		wg.Add(1)
		go f1(wg)
	}
	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 1000)
	wg.Done()
}


// // You can edit this code!
// // Click here and start typing.
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界", ls([]int{1, 1, 0, 0, 1, 0, 1, 0, 0, 1}))
// }

// // [1,1,0,0,1,0,1,0,0,1]

// func longestSubarray(nums []int) int {
// 	maxInterval := 0
// 	currentInterval := 0
// 	score := 1
// 	cnt := 0

// 	for i := 0; i < len(nums); i++ {

// 		if nums[i] == 1 {
// 			cnt++
// 			currentInterval++
// 			if i == len(nums)-1 {
// 				if currentInterval > maxInterval {
// 					maxInterval = currentInterval
// 					score = 1
// 				}
// 				break
// 			}
// 		} else if currentInterval > 0 && score > 0 {
// 			if i < len(nums)-1 && nums[i+1] == 1 {
// 				i++
// 				currentInterval++
// 				score--
// 			} else {
// 				if currentInterval > maxInterval {
// 					maxInterval = currentInterval
// 					currentInterval = 0
// 					score = 1
// 				}
// 			}
// 		} else {
// 			if currentInterval > maxInterval {
// 				maxInterval = currentInterval
// 				currentInterval = 0
// 				score = 1
// 			}
// 		}
// 	}
// 	if currentInterval > maxInterval {
// 		maxInterval = currentInterval
// 	}
// 	if cnt == len(nums) {
// 		return len(nums) - 1
// 	}
// 	return maxInterval
// }
