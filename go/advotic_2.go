package main

import (
	"fmt"
)

func checkSum(n int) (res []int, mod int, status bool) {
	result := make([]int, 0)
	var j int
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			j = i * 2
		} else {
			j = i
		}

		// if j > 9 {
		// 	r := strconv.Itoa(j)
		// 	fmt.Println(r)
		// }
		result = append(result, j)
	}

	total := 0

	for z := 0; z < len(result); z++ {
		total = total + result[z]
	}

	mod = 10 - (total % 10)

	status = true
	if mod > 0 {
		status = false
	}

	return result, mod, status
}

func main() {
	_, debt_mod, stat := checkSum(6)

	fmt.Println(debt_mod)
	fmt.Println(stat)
}
