package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	URL_DATABASE = ""
	PORT         = 0
)

func Config() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		PORT = 9000
	}

	URL_DATABASE = fmt.Sprint(os.Getenv("URL_DATABASE_ENV"))

}