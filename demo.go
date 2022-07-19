package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello, World!", time.Now())

	var sliceA []int
	var mapA = make(map[int]int)

	fmt.Println("sliceA:", sliceA, len(sliceA), cap(sliceA))
	fmt.Println("mapA:", mapA, len(mapA))
	if sliceA == nil {
		fmt.Println("sliceA is nil")
	}
	sliceA = append(sliceA, 1)
	fmt.Println("sliceA:", sliceA, len(sliceA), cap(sliceA))

	mapA[1] = 3
	fmt.Println("mapA:", mapA, len(mapA))

	sliceB := []int{1, 2, 3, 4, 5, 6}
	sliceC := sliceB[1:3]
	sliceC[0] = 3
	fmt.Println("sliceB:", sliceB, len(sliceB), cap(sliceB))
	fmt.Println("sliceC:", sliceC, len(sliceC), cap(sliceC))

}
