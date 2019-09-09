package main

import (
	"fmt"
)

func Nol(A []int) int {
	change:=0

	for i,v := range A {
		if i%2==0 && v==1 {
			change++
		} else if i%2==1 && v==0 {
				change++
		}
	}

	return change
}

func Satu(A []int) int {
	change:=0

	for i,v := range A {
		if i%2==0 && v==0 {
			change++
		} else if i%2==1 && v==1 {
			change++
		}
	}

	return change
}

func Koin(A []int) int {
	nol := Nol(A)
	satu := Satu(A)

	if int(nol) > int(satu) {
		return satu
	}

	return nol
}

func main() {
	answer:= Koin([]int{1,1,0,1,1})

	fmt.Println(answer)
}