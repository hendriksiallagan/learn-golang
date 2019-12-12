package main

import (
	"fmt"
	"strings"
)

func DifferentCases(str string) (result string) {
  
 f := func(r rune) bool {
        return r < 'A' || r > 'z'
    }

var splstr []string

result = ""

	spec_char := []string{" ", "!","@","#","$","%","^","&","*","(",")","<",">","/","?","-","_","+","=", ";", ":"}
	
    if strings.IndexFunc(str, f) != -1 {
        for i:=0; i<len(spec_char); i++ {
		str=strings.Replace(str, spec_char[i], " ", -1)
		
	}
	
	splstr = strings.Split(str, " ")
	
	for j:=0; j<len(splstr); j++ {
		result += strings.Title(strings.ToLower(splstr[j]))

	}
	
	
	
    }
	return result
}


func main() {
	test := DifferentCases("cats AND*Dogs-are Awesome")
	
	fmt.Println(test)
}
