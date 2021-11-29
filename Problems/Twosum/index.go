package main

import "fmt"

func main() {
	target := 9
	nums := []int{2, 7, 11, 15}

	a := twoSum(nums, target)
	fmt.Println(a)
}

func twoSum(nums []int, target int) []int {
	var setIndex []int
	var prevItem []int
	var answer []int

	for index, item1 := range nums {
		for j, item2 := range prevItem {
			if nums[index]+item2 == target {
				answer = append(answer, setIndex[j], index)
				return answer
			}
		}
		setIndex = append(setIndex, index)
		prevItem = append(prevItem, item1)
	}
	return answer
}
