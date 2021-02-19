package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type config struct {
	MAX_GOROUTINES int
	BASE_PATH      string
	DATA_PATH      string
	CACHE_PATH     string
}

func (c *config) setBasePath() {
	fmt.Println("Base path add")
	file, err := os.Getwd()
	if err != nil {
		file = ""
	}
	c.BASE_PATH = file
}

func (c *config) setDataPath() {
	fmt.Println("Data path add")
	c.DATA_PATH = filepath.Join(c.BASE_PATH, "data")
}

func (c *config) setCachePath() {
	fmt.Println("Data path add")
	c.CACHE_PATH = filepath.Join(c.BASE_PATH, "cache")
}

func (c *config) SetMaxGoroutines(maxGoroutines int) {
	fmt.Println("Max goroutines to ", maxGoroutines)
	c.MAX_GOROUTINES = maxGoroutines
}

func GetConfig() *config {
	conn := config{}
	conn.setBasePath()
	conn.setCachePath()
	conn.setDataPath()
	return &conn
}
