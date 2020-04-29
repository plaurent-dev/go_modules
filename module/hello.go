package hello

import (
    "fmt"
	"rsc.io/quote/v3"
)

func Person() string {
    fmt.Println("Personne !?")
    return "Personne !?"
}

func Hello() string {
    return quote.HelloV3()
}

func Proverb() string {
    return quote.Concurrency()
}