package qlikapi

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"errors"
	"fmt"
	"bytes"
)

type ReloadStruct struct {
	AppID    string     `json:"AppID"`
}

func Reload(apiKey string, tenantUrl string, appId string) (ReloadStruct, error) {

	var reload ReloadStruct

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	reload = ReloadStruct{
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


