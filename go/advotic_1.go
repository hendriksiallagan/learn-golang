package main

import "fmt"

func RotateToRight(n int, rot int) (r []int, temp int) {
	result := make([]int, 0)
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}

	//temp2 := 0
	for j := 0; j < len(result); j++ {
		temp = result[j]
		result[j] = result[len(result)-1]
		result[len(result)-1] = temp

	}

	return result, temp
}

func RotateToLeft(n int, rot int) (r []int, temp int) {
	result := make([]int, 0)
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}

	for j := 0; j < len(result); j++ {
		if j == 0 {
			temp = result[len(result)-1]
			result[len(result)-1] = result[j]
			result[j] = temp
		} else if j < len(result)-1 {
			temp = result[j-1]
			result[j-1] = result[j]
			result[j] = temp
		}
	}

	return result, temp
}

func main() {
	test, te2 := RotateToLeft(5, 1)

	fmt.Println(test)
	fmt.Println(te2)
}
