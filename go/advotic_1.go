package main

import "fmt"

func Rotate(n int, rot int) (r []int, temp int) {
	result := make([]int, 0)
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}

	temp2 := 0
	for j := 0; j < len(result); j++ {
		if j == 0 {
			temp = result[j+1]
			result[j+1] = result[j]
		} else if j < len(result)-1 {
			temp2 = result[j+1]
			result[j+1] = temp
			temp = temp2
		} else {
			result[0] = temp
		}
	}

	return result, temp
}

func main() {
	test, te2 := Rotate(5, 2)

	fmt.Println(test)
	fmt.Println(te2)
}
