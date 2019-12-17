package main

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"errors"
	"fmt"
	"bytes"
)

type Reload struct {
	AppID    string     `json:"AppID"`
}

func reload(apiKey string, tenantUrl string, appId string) (Reload, error) {

	var reload Reload

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	reload = Reload{
		AppID:         appId,
	}

	reloadJson, err := json.Marshal(reload)
	if err != nil {
		return reload, err
	}

	req, err := http.NewRequest("POST", tenantUrl + "/api/v1/reloads", bytes.NewBuffer(reloadJson))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return reload, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return reload, errors.New(fmt.Sprintf("Unexpected Status Code: %d", resp.StatusCode))
	}

	return reload, nil
}


