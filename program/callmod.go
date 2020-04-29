package main

import (
	"fmt"
	"github.com/plaurent-django/hello"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Println("Hello World !")
	fmt.Println(hello.Hello())
	fmt.Println(quote.Concurrency())
	fmt.Println(hello.Person())
}