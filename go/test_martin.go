package main

import (
	"fmt"
	"strings"
)

func Martin(a string, b string) string {
	var result string
	var awal int
	var akhir int
	var awStr string
	var akStr string

	result = ""

	lenB := len(b)
	lenA := len(a)

	if lenA>3*lenB || lenB>3*lenA {
		return "false"
	}

	if lenA>=lenB {
		awal = lenA
		akhir = lenB
		awStr = a
		akStr = b
	} else if lenA<lenB {
		awal = lenB
		akhir = lenA
		awStr = b
		akStr = a
	}

	s := strings.Split(awStr, "")
	t := strings.Split(akStr, "")

	q:=0
	w:=0
	for i:=0; i<awal; i++ {
		result += s[i]
		q++
		if q==2 && w<akhir{
			result += t[w]
			w++
			q=0
		}
	}


	return result
}

func main() {
	test := Martin("AAAAAA", "BB")

	fmt.Println(test)
}
