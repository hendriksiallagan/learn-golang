package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Structure struct {
	stuff []interface{}
}

func main() {
	url := "https://api.coinmarketcap.com/v1/ticker/?start=0&limit=100"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var decoded []interface{}
	err = json.Unmarshal(body, &decoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(decoded)
}
