package src

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrlucca/swapi_consumer/config"
)

func verifyConnectionInUrlBase() error {
	response, err := http.Get(config.URLS["BASE"])
	log.Println("Verify url base")
	if err != nil {
		return fmt.Errorf("Url base not found")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		error_msg := fmt.Sprintf("Response code is not good! -> %v", response.StatusCode)
		return fmt.Errorf(error_msg)
	}
	return nil

}
