package main

import (
	"fmt"
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

	appearancesMap := make(map[int]int)
	var result int
	for _, num := range left {
		value, ok := appearancesMap[num]
		if ok {
			result += value
			continue
		}

		var count int
		for _, r := range right {
			if r == num {
				count += 1
			}
		}

		appearancesMap[num] = count * num
		result += count * num
	}

	fmt.Println(result)
}
