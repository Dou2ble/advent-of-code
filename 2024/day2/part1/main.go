package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var data [][]int
	numOfSafeReports := 0

	for {
		var nums []int
		input, err := reader.ReadString('\n')
		if strings.TrimSpace(input) == "" || err != nil {
			break
		}

		numStrings := strings.Split(input, " ")
		for _, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			if err != nil {
				break
			}

			nums = append(nums, num)
		}

		data = append(data, nums)
	}

	for _, report := range data {
		var increasing bool
		fail := false

		for i := 0; i < len(report)-1; i++ {
			// find if this sequence is increasing or decreasing
			if i == 0 {
				if report[0] < report[1] {
					increasing = true
				} else if report[0] > report[1] {
					increasing = false
				} else {
					// it is not a sucess if it is the same
					fail = true
					break
				}
			}

			difference := report[i] - report[i+1]
			if difference == 0 {
				fail = true
				break
			} else if difference > 0 {
				if increasing {
					fail = true
					break
				}
			} else {
				if !increasing {
					fail = true
					break
				}
			}

			if int(math.Abs(float64(difference))) > 3 {
				fail = true
				break
			}
		}

		if !fail {
			numOfSafeReports++
		}
	}

	fmt.Println(numOfSafeReports)
}
