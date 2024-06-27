package main

import (
	"broker/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

var validatorInfoArray []*models.Info

func (app *Config) ValidatorHandler(w http.ResponseWriter, r *http.Request) {

	// validatorInfoArray := make([]*models.Info, len(app.validatorKeys))
	// validatorInfoArray := []*models.Info{}

	slog.Info("ValidatorHandler", "VALIDATOR KEYS=", len(app.validatorKeys))
	for i := 0; i < len(app.validatorKeys); i++ {
		// if i == 0 {
		// slog.Info("ValidatorHandler", "key", app.validatorKeys[i], "i", i)
		models, err := app.getInfoByKey(app.validatorKeys[i])

		if err != nil {
			slog.Error("getInfoByKey return nill for", "key", app.validatorKeys[i])
		}
		if models != nil {
			slog.Info("ValidatorHandler", "adding validator", app.validatorKeys[i], "i", i)
			validatorInfoArray = append(validatorInfoArray, models)
			//validatorInfoArray[i] = models
		}
	}
	app.writeJSON(w, http.StatusOK, validatorInfoArray)
}

func (app *Config) ValidatorStatusHandler(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	msg, err := app.CreateValidatorStatusResponse()
	payload.Message = msg
	if err != nil {
		payload.Error = true
		app.writeJSON(w, http.StatusBadRequest, payload)
	} else {
		app.writeJSON(w, http.StatusOK, payload)
	}
}
func (app *Config) CreateValidatorStatusResponse() (string, error) {
	if len(validatorInfoArray) != len(app.validatorKeys) {
		return string("Not all validators checked"), errors.New("Not all validators checked")
	}

	slog.Info("ValidatoStatusHandler", "VALIDATOR KEYS=", len(app.validatorKeys))
	for i := 0; i < len(app.validatorKeys); i++ {
		if validatorInfoArray[i].Status != "OK" {
			return string("Validator status is not OK "), errors.New("Validator status is not OK")
		}
		if validatorInfoArray[i].Data.Status != "active_online" {
			return string("Validator status is has to be active_online but  got" + validatorInfoArray[i].Data.Status), errors.New("Validator status is has to be active_online")
		}
	}

	ret := "Checked " + string(len(app.validatorKeys)) + " validators everything looks OK"
	return ret, nil
}

// GetCredentials implements Api.
func (app *Config) getInfoByKey(key string) (*models.Info, error) {

	// key := "0xa94ed867357ed9165a5ed10c10be9961b08430bf52eec53a0de6768f5b23c0077038d7ecb8da7cdfb6dc36d3816f830a"
	url := fmt.Sprintf("https://beaconcha.in/api/v1/validator/%s", key)

	slog.Info("url to auth manager", "url", url)

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)

	// req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed ot get user credentials %v", resp.StatusCode)
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// myString := string(b)
	// slog.Info("body from request", " is", myString)
	var creds *models.Info
	err = json.Unmarshal(b, &creds)
	if err != nil {
		return nil, err
	}

	// slog.Info("test", "unmarshal", creds)
	return creds, nil

}
