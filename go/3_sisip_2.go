package main

import (
	"fmt"
	"strconv"
	"strings"
)
//(B string, n int)

func Sisip2(A int) int {
	strA := strconv.Itoa(A)
	s := strings.Split(strA, "")

	//cf := ""
	idx := -1
	//fmt.Println(idx)

	for i,v := range s {
		intV,_ := strconv.Atoi(v)
		//if intV <= 5 {
		//	return strA, i
		//}
		if A>=0 && intV <= 5 {
			//return strA, i
			if idx==-1 {
				idx = i
			}
		} else if A<=0 && intV >= 5 {
			//return strA, i
			if idx == -1 {
				idx = i
			}
		}
		//} else {
		//	if idx==-1 {
		//		idx=len(s)
		//	}
		//	//idx=len(s)
		//}
	}

	if idx==-1 {
				idx=len(s)
	}
	//fmt.Println(strA)
	//fmt.Println(idx)
	//idx = len(s)
	q := strA[:idx] + "5" + strA[idx:]
	result,_:= strconv.Atoi(q)

	//return strA, len(s)

	return result
}


func main() {
	//C, idx := Sisip(789)
	//
	//q := C[:idx] + "5" + C[idx:]
	//
	//res,_:= strconv.Atoi(q)

	//res :=  Sisip2(-11)

	res :=  Sisip2(0)

	fmt.Println(res)
}