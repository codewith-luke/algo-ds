package main

import (
	"fmt"

	"github.com/lukeinthecloud/algo-ds/binary"
)

func main() {
	x := []int{10, 20, 30, 40, 50}
	pos, _ := binary.Search(x, 10)
	fmt.Println("position:", pos)
}
