package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName 	string `json:"Name"`
	Age 		int
	Address 	string
}

func main() {
	//json unmarshal ->json string diconvert menjadi objek/array
	var jsonString = `{"Name": "john wick", "Age": 27, "Address": "kampret"}`
	var jsonData = []byte(jsonString)

	var data User

	var errUnMars = json.Unmarshal(jsonData, &data)

	fmt.Println(errUnMars)
	fmt.Println(data)

	//json marshal -> objek/array diconvert menjadi json
	var object = []User{{"john wick", 27, "croot"}, {"edria", 20, "woow"}}
	var jsonMars, errMars = json.Marshal(object)

	var jsonStr = string(jsonMars)

	fmt.Println(jsonMars)
	fmt.Println(jsonStr)
	fmt.Println(errMars)
}
