package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mrlucca/swapi_consumer/config"
	"github.com/mrlucca/swapi_consumer/models"
)

func GetPeople() func(peopleReference string) (models.People, error) {
	log.Println("Init get people flow")
	err := verifyConnectionInUrlBase()
	if err != nil {
		panic(err.Error())
	}
	url_base := fmt.Sprintf("%v/people/", config.URLS["BASE"])
	return func(peopleReference string) (models.People, error) {
		url := url_base + peopleReference
		log.Println("Get people -> " + url)
		response, err := http.Get(url)
		if err != nil {
			msg_err := fmt.Sprintf("Not possible get page, status code = %v", response.StatusCode)
			log.Fatalln(msg_err)
		}
		log.Println("Trasnform people data in string")
		body, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			return models.People{}, fmt.Errorf("No users in stream of bytes")
		}
		var people models.People
		err = json.Unmarshal(body, &people)
		if err != nil {
			log.Fatalln("Object is not found")
		}

		return people, nil

	}
}
