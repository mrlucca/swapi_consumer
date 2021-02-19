package main

import (
	"fmt"
	"strconv"

	"github.com/mrlucca/swapi_consumer/src"
)

func main() {
	// spyderConfig := config.GetConfig()
	// fmt.Println(spyderConfig)
	// spyderConfig.SetMaxGoroutines(10)
	// fmt.Println(spyderConfig)
	getPeople := src.GetPeople()
	for i := 1; i < 50; i++ {
		peopleReference := strconv.Itoa(i)
		people, err := getPeople(peopleReference)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(people)
	}
}
