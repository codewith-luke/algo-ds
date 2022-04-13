package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30, 40, 50}
	pos, _ := Search(x, 20)
	fmt.Println(pos)
}

func Search(list []int, answer int) (int, error) {
	tail := 0
	head := len(list) - 1

	for tail <= head {
		mid := tail + head
		guess := list[mid]

		if answer == guess {
			return mid, nil
		} else if guess > answer {
			head = mid - 1
		} else {
			tail = mid + 1
		}
	}

	return -1, fmt.Errorf("not found")
}
