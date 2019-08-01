package main

import (
    "fmt"
    "strings"
"strconv"
)

func main() {
    s := strings.Split("OOXXOXXOOO", "")
    temp := 0

    result := ""
    
    for i,v := range s {
		if v == "O" {
		  temp++
		} else {
		  temp=0
		}

		temps := strconv.Itoa(temp)
		result += temps
		
		if i != len(s)-1 {
			result += "+"
		}
    }

    fmt.Println(result)
}