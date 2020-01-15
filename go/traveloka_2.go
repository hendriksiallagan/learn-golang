package main

import (
	"fmt"
	"reflect"
	"strings"
)

var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func inArray(val interface{}, array interface{}) (exists bool) {

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}

	return false
}

func getSpecialSubstring(s string, k int32, charValue string) int32 {
	splitString := strings.Split(charValue, "")
	result := make([]string, 0)
	for i := 0; i < 26; i++ {
		if splitString[i] == "0" {
			result = append(result, arr[i])
		}
	}

	x := 0
	y := 0

	ab := make([]int, 0)

	sp := strings.Split(s, "")
	for j := 0; j < len(sp); j++ {
		y++
		if inArray(sp[j], result) == true {
			x++
			if int(x) == int(k) {
				ab = append(ab, y)
			}
		}
	}

	max := int32(ab[0])

	for l := 1; l < len(ab); l++ {
		if int32(ab[l]) > max {
			max = int32(ab[l])
		}
	}

	return max
}

func main() {

	//test := getSpecialSubstring("giraffe", 2, "01111001111111111011111111")
	test := getSpecialSubstring("special", 1, "00000000000000000000000000")
	fmt.Println(test)
}
