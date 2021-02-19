package main

import (
	"fmt"

	"github.com/mrlucca/swapi_consumer/src"
)

func main() {
	done := make(chan bool, 1)
	go src.GetPeople(done)
	if <-done {
		fmt.Println("close")
	}
}
