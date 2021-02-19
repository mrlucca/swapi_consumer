package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/mrlucca/swapi_consumer/config"
	"github.com/mrlucca/swapi_consumer/models"
)

func getData(id string, jobs <-chan int, rawData chan<- []byte) {
	url := fmt.Sprintf("%v/people/%v", config.URLS["BASE"], id)
	log.Println("Get people -> " + url)
	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		msg_err := fmt.Sprintf("Not possible get page, status code = %v", response.StatusCode)
		log.Fatalln(msg_err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	rawData <- body

}

func formatData(body []byte) {
	log.Println("Trasnform people data in string")
	var people models.People
	err := json.Unmarshal(body, &people)
	if err != nil {
		log.Fatalln("Object is not found")
	}
	fmt.Println(people)
}

func GetPeople(done chan<- bool) {
	const numOfGet int = 80
	jobs := make(chan int, numOfGet)
	rawData := make(chan []byte, numOfGet)
	for i := 1; i <= numOfGet; i++ {
		id := strconv.Itoa(i)
		go getData(id, jobs, rawData)
	}
	for i := 1; i <= numOfGet; i++ {
		jobs <- i
	}
	for i := 1; i <= numOfGet; i++ {
		response := <-rawData
		formatData(response)
	}
	close(jobs)
	done <- true

}
