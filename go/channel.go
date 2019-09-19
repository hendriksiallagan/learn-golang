package main

import
(
	"fmt"
	"runtime"
)


func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string) //channel messages

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data //proses pengiriman dari variable data ke channel messages
	}

	go sayHelloTo("john wick")
	go sayHelloTo("edria")
	go sayHelloTo("aaa")

	var messages1 = <-messages
	fmt.Println(messages1)

	var messages2 = <-messages
	fmt.Println(messages2)

	test:="fuck"

	fmt.Printf("hai %s", test)
	//fmt.Sprintf("hai %s", test)

	var messages3 = <-messages
	fmt.Println(messages3)
}