package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Sisip(A int) (B string, n int) {
	strA := strconv.Itoa(A)
	s := strings.Split(strA, "")

	for i,v := range s {
		intV,_ := strconv.Atoi(v)

		if A>0 && intV <= 5 {
			return strA, i
		} else if A<0 && intV >= 5 {
			return strA, i
		}
	}

	return strA, len(s)
}


func main() {
	C, idx := Sisip(-89)

	fmt.Println(C)
	fmt.Println(idx)

	q := C[:idx] + "5" + C[idx:]

	res,_:= strconv.Atoi(q)

	fmt.Println(res)
}