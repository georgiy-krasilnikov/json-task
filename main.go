package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	mod "json_task/models"
)

func main() {
	resp, err := http.Get("https://api.coinpaprika.com/v1/coins/btc-bitcoin")
	if err != nil {
		log.WithError(err).Fatal("there was an HTTP protocol error")
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Fatal("failed to read http.response body")
		os.Exit(1)
	}

	var message mod.Message
	if err := json.Unmarshal(body, &message); err != nil {
		log.WithError(err).Fatal("failed to unmarchal []bytes from http.response body")
		os.Exit(1)
	}

	fmt.Printf("%+v\n", message)
}
