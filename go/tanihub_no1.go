package main

import (
	"fmt"
	"strconv"
	"strings"
)

func TwoSum(arr []int) string {
	first := arr[0]

	var result string

	result = ""

	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if i != j {
				if arr[i]+arr[j] == first {
					if strings.Contains(result, strconv.Itoa(arr[i])) == false && strings.Contains(result, strconv.Itoa(arr[j])) == false {
						result += strconv.Itoa(arr[i]) + "," + strconv.Itoa(arr[j]) + " "
					}
				}
			}
		}
	}

	if result == "" {
		result = "-1"
	}

	return result
}

func main() {
	test := TwoSum([]int{17, 4, 5, 6, 10, 11, 4, -3, -5, 3, 15, 2, 7})

	fmt.Println(test)

}
