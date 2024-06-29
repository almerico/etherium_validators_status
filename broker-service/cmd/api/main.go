package main

import (
	"broker/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const webPort = "8081"

type Config struct {
	validatorKeys      []string
	validatorInfoArray [11]*models.Info
}

// var validatorKeys []string

func main() {
	dat, err := os.ReadFile("all_key.txt")

	if err != nil {
		log.Panic(err)
	}
	// fmt.Print(string(dat))

	// var validatorKeys []string
	app := Config{}

	err = json.Unmarshal(dat, &app.validatorKeys)
	fmt.Printf("Loaded %d keys\n", len(app.validatorKeys))
	fmt.Println(app.validatorKeys)

	if err != nil {
		log.Panic(err)
	}

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
