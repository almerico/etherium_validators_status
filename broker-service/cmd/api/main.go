package main

import (
	"broker/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron"
)

const webPort = "8081"

type Config struct {
	validatorKeys      []string
	validatorInfoArray [11]*models.Info
}

func main() {
	dat, err := os.ReadFile("all_key.txt")

	if err != nil {
		log.Panic(err)
	}
	app := Config{}

	err = json.Unmarshal(dat, &app.validatorKeys)
	fmt.Printf("Loaded %d keys\n", len(app.validatorKeys))
	fmt.Println(app.validatorKeys)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Starting broker service on port %s\n", webPort)
	c := cron.New()
	// c.AddFunc("*/5 * * * * *", func() { fmt.Println("Testing every 5 seconds.") })
	c.AddFunc("*/60 * * * * *", app.CheckValidatorsJob)
	c.Start()

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
