package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type cr struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

type people struct {
	Number int  `json:"number"`
	People []cr `json:"people"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//croot, _ := json.Marshal(&people1.People)

	// for _, v := range people1.People {
	// 	fmt.Println(v)
	// }

	//fmt.Println(people1.Number)
	fmt.Println(people1)
	fmt.Println(people1.People)

	for _, v := range people1.People {
		fmt.Println(v.Name)
	}
}
