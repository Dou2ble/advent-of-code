package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var left []int
	var right []int
	for {
		var leftInt, rightInt int

		_, err := fmt.Scanf("%d\t%d", &leftInt, &rightInt)
		if err != nil {
			break
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	sort.Ints(left)
	sort.Ints(right)

	var totalDistance int
	for i := range left {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(totalDistance)
}
