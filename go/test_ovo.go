package main

import (
	"fmt"
)

func revOvo(loop int, word string) (a int, b string) {
	var res string
	res=""

	for i:=len(word)-1; i>=0; i-- {
		res += fmt.Sprintf("%c", word[i])
	}

	var rescue string

	if loop > 1 {
		for j:=1; j<=loop; j++ {
			rescue = rescue + res
		}
	} else {
		rescue = res
	}





	return len(word), rescue
}

func main() {
	_, testb := revOvo(3, "iqbal martin")

	//fmt.Println(test)
	fmt.Println(testb)
}