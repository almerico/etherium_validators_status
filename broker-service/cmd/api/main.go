package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const webPort = "80"

type Config struct {
	validatorKeys []string
}

var validatorKeys []string

func main() {
	dat, err := os.ReadFile("all_key.txt")

	if err != nil {
		log.Panic(err)
	}
	fmt.Print(string(dat))

	// var validatorKeys []string
	err = json.Unmarshal(dat, &validatorKeys)
	fmt.Print(validatorKeys)

	if err != nil {
		log.Panic(err)
	}
	app := Config{}
	app.validatorKeys = validatorKeys

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
