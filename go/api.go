package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://ronreiter-meme-generator.p.rapidapi.com/fonts"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("x-rapidapi-host", "ronreiter-meme-generator.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "301073fbd7mshb6d892c3ef9efb2p1af3e9jsnb3e9ba337819")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
