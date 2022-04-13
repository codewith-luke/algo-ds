package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30, 40, 50}
	pos, _ := Search(x, 10)
	fmt.Println("position:", pos)
}

func Search(list []int, item int) (int, error) {
	low := 0
	high := len(list) - 1
	count := 0

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		count++

		fmt.Println("attempt:", guess, mid)

		if guess == item {
			fmt.Println("guesses:", count)
			return mid, nil
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, fmt.Errorf("not found")
}
