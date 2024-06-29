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

func (app *Config) getValidatorStatus() {
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
			app.validatorInfoArray[i] = models
			//validatorInfoArray[i] = models
		}
	}
}

func (app *Config) ValidatorHandler(w http.ResponseWriter, r *http.Request) {

	// validatorInfoArray := make([]*models.Info, len(app.validatorKeys))
	// validatorInfoArray := []*models.Info{}

	// validatorInfoArray = validatorInfoArray[:0]
	app.getValidatorStatus()
	app.writeJSON(w, http.StatusOK, app.validatorInfoArray)
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
		app.writeJSON(w, http.StatusBadRequest, err)
	} else {
		app.writeJSON(w, http.StatusOK, payload)
	}
}
func (app *Config) CreateValidatorStatusResponse() (string, error) {
	if len(app.validatorInfoArray) != len(app.validatorKeys) {
		msg := "validators checked " + string(len(app.validatorInfoArray)) + " validator in initial list" + string(len(app.validatorKeys))
		return msg, errors.New("")
	}

	for i := 0; i < 10; i++ {
		if len(app.validatorInfoArray) == len(app.validatorKeys) {
			break
		}
		time.Sleep(1 * time.Second)
	}
	slog.Info("ValidatorStatusHandler", "VALIDATOR KEYS=", len(app.validatorKeys))
	for i := 0; i < len(app.validatorKeys); i++ {
		if app.validatorInfoArray[i].Status != "OK" {
			return string("Validator status is not OK for " + app.validatorInfoArray[i].Data.Pubkey), errors.New("Exception")
		}
		if app.validatorInfoArray[i].Data.Status != "active_online" {
			return string("Validator status is has to be active_online but  got" + app.validatorInfoArray[i].Data.Status + " for " + app.validatorInfoArray[i].Data.Pubkey), errors.New("Exception")
		}
	}
	t1 := time.Now()
	ret := "Validators are ACTIVE checked at " + t1.String()
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
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed ot get user credentials %v", resp.StatusCode)
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var creds *models.Info
	err = json.Unmarshal(b, &creds)
	if err != nil {
		return nil, err
	}

	slog.Info("tegetInfoByKeyt", "unmarshal", creds)
	return creds, nil

}
