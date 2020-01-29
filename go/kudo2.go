package main

import (
	"bytes"
	"fmt"
	"net/http"

	"moul.io/http2curl"
)

func main() {
	data := bytes.NewBufferString(`{"hello":"world","answer":42}`)
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", data)
	req.Header.Set("Content-Type", "application/json")

	command, _ := http2curl.GetCurlCommand(req)
	fmt.Println(command)
}
