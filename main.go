package main

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	if value, ok := os.LookupEnv("LOG_LEVEL"); ok {
		level, err := logrus.ParseLevel(value)
		if err != nil {
			log.Fatal(err)
		}
		logrus.SetLevel(level)
	}

	var config config
	err := config.read()
	if err != nil {
		log.Fatal(err)
	}

	var output output
	var results []result

	status, err := config.HashicorpTerraformCloud.checkStatus()
	if err != nil {
		log.Fatal(err)
	}

	results = append(results, status...)

	output.results = results

	output.writeTable()
}
